package crc16

import (
	"fmt"
	"testing"
)

func TestChecksum(t *testing.T) {
	fmt.Println(Checksum([]byte("k1")) & 16383)
}

func BenchmarkBitAnd(b *testing.B) {
	var u uint16
	for i := 0; i < b.N; i++ {
		u = Checksum([]byte("k1")) & 16383
	}

	fmt.Println(u)
}

func BenchmarkMod(b *testing.B) {
	var u uint16
	for i := 0; i < b.N; i++ {
		u = Checksum([]byte("k1")) % 16384
	}

	fmt.Println(u)
}
