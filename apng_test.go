package adventofcode2016

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	"testing"
)

func fixedDelay(n, d uint16) DelayFunc { return func(int) (uint16, uint16) { return n, d } }

func TestEncoder_BoxesProgressive(t *testing.T) {
	buf := &memRW{}
	e, err := NewEncoder(buf, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 24; i++ {
		if err := e.Push(drawChecker(128, 64, i)); err != nil {
			t.Fatal(err)
		}
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
	b := buf.Bytes()
	if len(b) < 8 || !bytes.Equal(b[:8], pngSignature) {
		t.Fatalf("bad signature")
	}
	var (
		seenIHDR, seenIEND, seenAcTL  bool
		fcTLs, idat0, fdats, seqCount int
		numFrames, numPlays           uint32
		lastSeq                       int64 = -1
	)
	off := 8
	for off+8 <= len(b) {
		ln := int(binary.BigEndian.Uint32(b[off : off+4]))
		if off+8+ln+4 > len(b) {
			t.Fatalf("truncated chunk")
		}
		typ := string(b[off+4 : off+8])
		data := b[off+8 : off+8+ln]
		switch typ {
		case ctIHDR:
			seenIHDR = true
		case ctacTL:
			seenAcTL = true
			numFrames = binary.BigEndian.Uint32(data[0:4])
			numPlays = binary.BigEndian.Uint32(data[4:8])
		case ctfcTL:
			fcTLs++
			seq := int64(binary.BigEndian.Uint32(data[0:4]))
			if seq <= lastSeq {
				t.Fatalf("non-increasing seq")
			}
			lastSeq = seq
			seqCount++
		case ctIDAT:
			idat0++
		case ctfdAT:
			fdats++
			seq := int64(binary.BigEndian.Uint32(data[0:4]))
			if seq <= lastSeq {
				t.Fatalf("non-increasing seq")
			}
			lastSeq = seq
			seqCount++
		case ctIEND:
			seenIEND = true
		}
		off += 8 + ln + 4
		if typ == ctIEND {
			break
		}
	}
	if !seenIHDR || !seenAcTL || !seenIEND {
		t.Fatalf("missing critical chunk")
	}
	if numFrames != 24 || fcTLs != 24 {
		t.Fatalf("frames=%d fcTL=%d", numFrames, fcTLs)
	}
	if numPlays != 1 {
		t.Fatalf("numPlays=%d want 1", numPlays)
	}
	if seqCount != fcTLs+fdats {
		t.Fatalf("seq mismatch: got %d want %d", seqCount, fcTLs+fdats)
	}
	if idat0 == 0 {
		t.Fatalf("no IDAT for frame0")
	}
}

func TestEncoder_GridFixed1s(t *testing.T) {
	buf := &memRW{}
	e, err := NewEncoder(buf, 0, fixedDelay(100, 100))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		img := drawChecker(64, 64, i)
		if err := e.Push(img); err != nil {
			t.Fatal(err)
		}
	}
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
	b := buf.Bytes()
	off := 8
	var frames, plays uint32
	for off+8 <= len(b) {
		ln := int(binary.BigEndian.Uint32(b[off : off+4]))
		if off+8+ln+4 > len(b) {
			t.Fatalf("trunc")
		}
		typ := string(b[off+4 : off+8])
		if typ == ctacTL {
			frames = binary.BigEndian.Uint32(b[off+8 : off+12])
			plays = binary.BigEndian.Uint32(b[off+12 : off+16])
			break
		}
		off += 8 + ln + 4
	}
	if frames != 10 {
		t.Fatalf("frames=%d want 10", frames)
	}
	if plays != 0 {
		t.Fatalf("num_plays=%d want 0", plays)
	}
}

func TestEncoder_NoFrames_Error(t *testing.T) {
	buf := &memRW{}
	e, err := NewEncoder(buf, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := e.Close(); err == nil {
		t.Fatalf("expected error on Close with no frames")
	}
}

func drawChecker(w, h, i int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	sz := 8
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v := ((x/sz + y/sz + i) % 2) * 255
			img.Set(x, y, color.NRGBA{uint8(v), uint8(v), uint8(v), 255})
		}
	}
	return img
}
