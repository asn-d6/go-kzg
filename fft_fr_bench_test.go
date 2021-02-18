package kzg

import (
	"fmt"
	"github.com/protolambda/go-kzg/bls"
	"testing"
)

func benchFFT(scale uint8, b *testing.B) {
	fs := NewFFTSettings(scale)
	data := make([]bls.Big, fs.maxWidth, fs.maxWidth)
	for i := uint64(0); i < fs.maxWidth; i++ {
		bls.CopyBigNum(&data[i], bls.RandomBig())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out, err := fs.FFT(data, false)
		if err != nil {
			b.Fatal(err)
		}
		if len(out) != len(data) {
			panic("output len doesn't match input")
		}
	}
}

func BenchmarkFFTSettings_FFT(b *testing.B) {
	for scale := uint8(4); scale < 16; scale++ {
		b.Run(fmt.Sprintf("scale_%d", scale), func(b *testing.B) {
			benchFFT(scale, b)
		})
	}
}
