package hsalsa

import (
	"bytes"
	"testing"
)

// TestHSalsa20
func TestHSalsa20(t *testing.T) {
	var (
		shared = [32]byte{
			0x4a, 0x5d, 0x9d, 0x5b, 0xa4, 0xce, 0x2d, 0xe1,
			0x72, 0x8e, 0x3b, 0xf4, 0x80, 0x35, 0x0f, 0x25,
			0xe0, 0x7e, 0x21, 0xc9, 0x47, 0xd1, 0x9e, 0x33,
			0x76, 0xf0, 0x9b, 0x3c, 0x1e, 0x16, 0x17, 0x42,
		}

		zero = [32]byte{}

		c = [16]byte{
			0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
			0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b,
		}

		firstKey = make([]byte, 32)

		expect = []byte{
			0x1b, 0x27, 0x55, 0x64, 0x73, 0xe9, 0x85, 0xd4,
			0x62, 0xcd, 0x51, 0x19, 0x7a, 0x9a, 0x46, 0xc7,
			0x60, 0x09, 0x54, 0x9e, 0xac, 0x64, 0x74, 0xf2,
			0x06, 0xc4, 0xee, 0x08, 0x44, 0xf6, 0x83, 0x89,
		}
	)

	firstKey = HSalsa20(firstKey[:0], zero[:], shared[:], c[:])

	if !bytes.Equal(expect, firstKey) {
		t.Error("expected result did not match computed", expect, firstKey)
	}
}

// TestHSalsa202
func TestHSalsa202(t *testing.T) {
	var (
		firstKey = [32]byte{
			0x1b, 0x27, 0x55, 0x64, 0x73, 0xe9, 0x85, 0xd4,
			0x62, 0xcd, 0x51, 0x19, 0x7a, 0x9a, 0x46, 0xc7,
			0x60, 0x09, 0x54, 0x9e, 0xac, 0x64, 0x74, 0xf2,
			0x06, 0xc4, 0xee, 0x08, 0x44, 0xf6, 0x83, 0x89,
		}

		noncePrefix = [16]byte{
			0x69, 0x69, 0x6e, 0xe9, 0x55, 0xb6, 0x2b, 0x73,
			0xcd, 0x62, 0xbd, 0xa8, 0x75, 0xfc, 0x73, 0xd6,
		}

		c = [16]byte{
			0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
			0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b,
		}

		secondKey = make([]byte, 32)

		expect = []byte{
			0xdc, 0x90, 0x8d, 0xda, 0x0b, 0x93, 0x44, 0xa9,
			0x53, 0x62, 0x9b, 0x73, 0x38, 0x20, 0x77, 0x88,
			0x80, 0xf3, 0xce, 0xb4, 0x21, 0xbb, 0x61, 0xb9,
			0x1c, 0xbd, 0x4c, 0x3e, 0x66, 0x25, 0x6c, 0xe4,
		}
	)

	secondKey = HSalsa20(secondKey[:0], noncePrefix[:], firstKey[:], c[:])

	if !bytes.Equal(expect, secondKey) {
		t.Error("expected result did not match computed", expect, secondKey)
	}
}

// TestHSalsa203
func TestHSalsa203(t *testing.T) {
	var (
		k = [32]byte{
			0xee, 0x30, 0x4f, 0xca, 0x27, 0x00, 0x8d, 0x8c,
			0x12, 0x6f, 0x90, 0x02, 0x79, 0x01, 0xd8, 0x0f,
			0x7f, 0x1d, 0x8b, 0x8d, 0xc9, 0x36, 0xcf, 0x3b,
			0x9f, 0x81, 0x96, 0x92, 0x82, 0x7e, 0x57, 0x77,
		}

		in = [16]byte{
			0x81, 0x91, 0x8e, 0xf2, 0xa5, 0xe0, 0xda, 0x9b,
			0x3e, 0x90, 0x60, 0x52, 0x1e, 0x4b, 0xb3, 0x52,
		}

		c = [16]byte{
			0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
			0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b,
		}

		out = make([]byte, 32)

		expect = []byte{
			0xbc, 0x1b, 0x30, 0xfc, 0x07, 0x2c, 0xc1, 0x40,
			0x75, 0xe4, 0xba, 0xa7, 0x31, 0xb5, 0xa8, 0x45,
			0xea, 0x9b, 0x11, 0xe9, 0xa5, 0x19, 0x1f, 0x94,
			0xe1, 0x8c, 0xba, 0x8f, 0xd8, 0x21, 0xa7, 0xcd,
		}
	)

	out = HSalsa20(out[:0], in[:], k[:], c[:])

	if !bytes.Equal(expect, out) {
		t.Error("expected result did not match computed", expect, out)
	}
}

// BenchmarkHSalsa20
func BenchmarkHSalsa20(b *testing.B) {
	var (
		shared = [32]byte{
			0x4a, 0x5d, 0x9d, 0x5b, 0xa4, 0xce, 0x2d, 0xe1,
			0x72, 0x8e, 0x3b, 0xf4, 0x80, 0x35, 0x0f, 0x25,
			0xe0, 0x7e, 0x21, 0xc9, 0x47, 0xd1, 0x9e, 0x33,
			0x76, 0xf0, 0x9b, 0x3c, 0x1e, 0x16, 0x17, 0x42,
		}

		zero = [32]byte{}

		c = [16]byte{
			0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x20, 0x33,
			0x32, 0x2d, 0x62, 0x79, 0x74, 0x65, 0x20, 0x6b,
		}

		firstKey = make([]byte, 16)
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		firstKey = HSalsa20(firstKey[:0], zero[:], shared[:], c[:])
	}
}
