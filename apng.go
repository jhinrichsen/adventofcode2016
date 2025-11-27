package adventofcode2016

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"image"
	"image/draw"
	"image/png"
	"io"
	"math"
	"os"
)

type DelayFunc func(i int) (num, den uint16)

func DefaultDelay(i int) (num, den uint16) {
	if i < 3 {
		return 30, 100
	}
	if i < 10 {
		return 15, 100
	}
	return 8, 100
}

type Encoder struct {
	w             io.ReadWriteSeeker
	ownsFile      bool
	f             *os.File
	loops         uint64
	delay         DelayFunc
	started       bool
	closed        bool
	nFrames       uint32
	seq           uint32
	width, height uint32
	acTLOffset    int64
	wroteIEND     bool
}

func NewEncoder(w io.ReadWriteSeeker, loops uint64, delay DelayFunc) (*Encoder, error) {
	if w == nil {
		return nil, errors.New("apng: nil writer")
	}
	if delay == nil {
		delay = DefaultDelay
	}
	return &Encoder{w: w, loops: loops, delay: delay}, nil
}

func NewFileEncoder(filename string, loops uint64, delay DelayFunc) (*Encoder, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	e, err := NewEncoder(f, loops, delay)
	if err != nil {
		_ = f.Close()
		return nil, err
	}
	e.ownsFile, e.f = true, f
	return e, nil
}

func (e *Encoder) Push(img image.Image) error {
	if e.closed {
		return errors.New("apng: encoder closed")
	}
	if img == nil {
		return errors.New("apng: nil image")
	}
	norm := toNRGBA(img)
	_, ihdr, idats, err := encodeAndSlice(norm)
	if err != nil {
		return err
	}
	w := binary.BigEndian.Uint32(ihdr[0:4])
	h := binary.BigEndian.Uint32(ihdr[4:8])
	if !e.started {
		if err := writeBytes(e.w, pngSignature); err != nil {
			return err
		}
		if err := writeChunk(e.w, ctIHDR, ihdr); err != nil {
			return err
		}
		off, _ := e.w.Seek(0, io.SeekCurrent)
		e.acTLOffset = off
		if err := writeAcTL(e.w, 0, mapLoopsToNumPlays(e.loops)); err != nil {
			return err
		}
		e.started = true
		e.width, e.height = w, h
	} else if w != e.width || h != e.height {
		return errors.New("apng: frame dimensions mismatch")
	}
	num, den := e.delay(int(e.nFrames))
	if den == 0 {
		den = 100
	}
	if err := writeFcTL(e.w, &e.seq, w, h, 0, 0, num, den, 0, 0); err != nil {
		return err
	}
	if e.nFrames == 0 {
		for _, p := range idats {
			if err := writeChunk(e.w, ctIDAT, p); err != nil {
				return err
			}
		}
	} else {
		for _, p := range idats {
			if err := writeFdAT(e.w, &e.seq, p); err != nil {
				return err
			}
		}
	}
	e.nFrames++
	return nil
}

func (e *Encoder) Close() error {
	if e.closed {
		return nil
	}
	defer func() { e.closed = true }()
	if e.nFrames == 0 {
		return errors.New("apng: no frames pushed")
	}
	if !e.wroteIEND {
		if err := writeChunk(e.w, ctIEND, nil); err != nil {
			return err
		}
		e.wroteIEND = true
	}
	if _, err := e.w.Seek(e.acTLOffset, io.SeekStart); err != nil {
		return err
	}
	if err := writeAcTL(e.w, e.nFrames, mapLoopsToNumPlays(e.loops)); err != nil {
		return err
	}
	if _, err := e.w.Seek(0, io.SeekEnd); err != nil {
		return err
	}
	if e.ownsFile && e.f != nil {
		return e.f.Close()
	}
	return nil
}

type Seq[T any] func(yield func(T) bool)

type NextFunc func() (image.Image, bool)

func PackAPNGToMem(seq Seq[image.Image], loops uint64, delay DelayFunc) ([]byte, error) {
	m := &memRW{}
	if err := PackAPNGFromSeq(seq, loops, m, delay); err != nil {
		return nil, err
	}
	return m.Bytes(), nil
}

func PackAPNGFromSeq(seq Seq[image.Image], loops uint64, w io.ReadWriteSeeker, delay DelayFunc) error {
	e, err := NewEncoder(w, loops, delay)
	if err != nil {
		return err
	}
	seq(func(img image.Image) bool {
		if err2 := e.Push(img); err == nil && err2 != nil {
			err = err2
		}
		return err == nil
	})
	if err != nil {
		return err
	}
	return e.Close()
}

func PackAPNGFileFromSeq(seq Seq[image.Image], loops uint64, filename string, delay DelayFunc) error {
	e, err := NewFileEncoder(filename, loops, delay)
	if err != nil {
		return err
	}
	seq(func(img image.Image) bool {
		if err2 := e.Push(img); err == nil && err2 != nil {
			err = err2
		}
		return err == nil
	})
	if err != nil {
		_ = e.Close()
		return err
	}
	return e.Close()
}

var pngSignature = []byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}

const (
	ctIHDR = "IHDR"
	ctacTL = "acTL"
	ctfcTL = "fcTL"
	ctfdAT = "fdAT"
	ctIDAT = "IDAT"
	ctIEND = "IEND"
)

func mapLoopsToNumPlays(loops uint64) uint32 {
	if loops == math.MaxUint64 {
		return 0
	}
	if loops > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(loops)
}

func toNRGBA(src image.Image) *image.NRGBA {
	if dst, ok := src.(*image.NRGBA); ok {
		return dst
	}
	r := src.Bounds()
	dst := image.NewNRGBA(r)
	draw.Draw(dst, r, src, r.Min, draw.Src)
	return dst
}

func encodeAndSlice(img image.Image) (full []byte, ihdr []byte, idats [][]byte, err error) {
	var buf bytes.Buffer
	if err = png.Encode(&buf, img); err != nil {
		return
	}
	full = buf.Bytes()
	p := full
	if len(p) < 8 || !bytes.Equal(p[:8], pngSignature) {
		err = errors.New("apng: bad PNG signature from encoder")
		return
	}
	off := 8
	for off+8 <= len(p) {
		ln := int(binary.BigEndian.Uint32(p[off : off+4]))
		if off+8+ln+4 > len(p) {
			err = errors.New("apng: truncated chunk")
			return
		}
		typ := string(p[off+4 : off+8])
		data := p[off+8 : off+8+ln]
		switch typ {
		case ctIHDR:
			ihdr = append([]byte(nil), data...)
		case ctIDAT:
			cp := make([]byte, len(data))
			copy(cp, data)
			idats = append(idats, cp)
		}
		off += 8 + ln + 4
		if typ == ctIEND {
			break
		}
	}
	if len(ihdr) != 13 {
		err = errors.New("apng: IHDR not found or invalid")
	}
	if len(idats) == 0 {
		err = errors.New("apng: no IDAT chunks")
	}
	return
}

func writeChunk(w io.Writer, typ string, data []byte) error {
	var lenb [4]byte
	binary.BigEndian.PutUint32(lenb[:], uint32(len(data)))
	if _, err := w.Write(lenb[:]); err != nil {
		return err
	}
	if _, err := io.WriteString(w, typ); err != nil {
		return err
	}
	if len(data) > 0 {
		if _, err := w.Write(data); err != nil {
			return err
		}
	}
	binary.BigEndian.PutUint32(lenb[:], crc32Of(typ, data))
	_, err := w.Write(lenb[:])
	return err
}

func writeBytes(w io.Writer, b []byte) error { _, err := w.Write(b); return err }

func writeAcTL(w io.Writer, numFrames, numPlays uint32) error {
	var d [8]byte
	binary.BigEndian.PutUint32(d[0:4], numFrames)
	binary.BigEndian.PutUint32(d[4:8], numPlays)
	return writeChunk(w, ctacTL, d[:])
}

func writeFcTL(w io.Writer, seq *uint32, wpx, hpx, xo, yo uint32, dnum, dden uint16, dispose, blend uint8) error {
	var d [26]byte
	s := *seq
	binary.BigEndian.PutUint32(d[0:4], s)
	s++
	*seq = s
	binary.BigEndian.PutUint32(d[4:8], wpx)
	binary.BigEndian.PutUint32(d[8:12], hpx)
	binary.BigEndian.PutUint32(d[12:16], xo)
	binary.BigEndian.PutUint32(d[16:20], yo)
	binary.BigEndian.PutUint16(d[20:22], dnum)
	binary.BigEndian.PutUint16(d[22:24], dden)
	d[24] = dispose
	d[25] = blend
	return writeChunk(w, ctfcTL, d[:])
}

func writeFdAT(w io.Writer, seq *uint32, idatPayload []byte) error {
	var hdr [4]byte
	s := *seq
	binary.BigEndian.PutUint32(hdr[:], s)
	s++
	*seq = s
	buf := make([]byte, 4+len(idatPayload))
	copy(buf[0:4], hdr[:])
	copy(buf[4:], idatPayload)
	return writeChunk(w, ctfdAT, buf)
}

type memRW struct {
	b   []byte
	pos int64
}

func (m *memRW) Read(p []byte) (int, error) {
	if m.pos >= int64(len(m.b)) {
		return 0, io.EOF
	}
	n := copy(p, m.b[m.pos:])
	m.pos += int64(n)
	return n, nil
}
func (m *memRW) Write(p []byte) (int, error) {
	end := m.pos + int64(len(p))
	if end > int64(len(m.b)) {
		nb := make([]byte, end)
		copy(nb, m.b)
		m.b = nb
	}
	copy(m.b[m.pos:end], p)
	m.pos = end
	return len(p), nil
}
func (m *memRW) Seek(off int64, whence int) (int64, error) {
	np := m.pos
	if whence == io.SeekStart {
		np = off
	}
	if whence == io.SeekCurrent {
		np = m.pos + off
	}
	if whence == io.SeekEnd {
		np = int64(len(m.b)) + off
	}
	if np < 0 {
		return 0, errors.New("memRW: negative position")
	}
	m.pos = np
	return np, nil
}
func (m *memRW) Bytes() []byte { return m.b }

func crc32Of(typ string, data []byte) uint32 {
	h := crc32.NewIEEE()
	_, _ = io.WriteString(h, typ)
	_, _ = h.Write(data)
	return h.Sum32()
}
