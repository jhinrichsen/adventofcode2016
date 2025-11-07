//go:build !wasm
// +build !wasm

package adventofcode2016

import (
	"image"
	"io"
)

func PackAPNGFromChan(ch <-chan image.Image, loops uint64, w io.ReadWriteSeeker, delay DelayFunc) error {
	e, err := NewEncoder(w, loops, delay)
	if err != nil {
		return err
	}
	for img := range ch {
		if err := e.Push(img); err != nil {
			_ = e.Close()
			return err
		}
	}
	return e.Close()
}

func PackAPNGFileFromChan(ch <-chan image.Image, loops uint64, filename string, delay DelayFunc) error {
	e, err := NewFileEncoder(filename, loops, delay)
	if err != nil {
		return err
	}
	for img := range ch {
		if err := e.Push(img); err != nil {
			_ = e.Close()
			return err
		}
	}
	return e.Close()
}
