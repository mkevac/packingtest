package pt_test

import (
	"testing"

	"github.com/mkevac/packingtest"
)

func BenchmarkLittle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = pt.KeyLittleEndian()
	}
}

func BenchmarkBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = pt.KeyBigEndian()
	}
}
