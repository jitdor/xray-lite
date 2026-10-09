package quic

import (
	"bytes"
	"io"
	"testing"
)

// Examples from RFC 9000, Appendix A.1.
func TestReadQUICVarint(t *testing.T) {
	tests := []struct {
		in   []byte
		want uint64
	}{
		{[]byte{0xc2, 0x19, 0x7c, 0x5e, 0xff, 0x14, 0xe8, 0x8c}, 151288809941952652},
		{[]byte{0x9d, 0x7f, 0x3e, 0x7d}, 494878333},
		{[]byte{0x7b, 0xbd}, 15293},
		{[]byte{0x25}, 37},
		{[]byte{0x40, 0x25}, 37},
	}
	for _, tt := range tests {
		got, err := readQUICVarint(bytes.NewReader(tt.in))
		if err != nil || got != tt.want {
			t.Errorf("readQUICVarint(%x) = %d, %v; want %d", tt.in, got, err, tt.want)
		}
	}
	// truncated input must fail, not return a partial value
	if _, err := readQUICVarint(bytes.NewReader([]byte{0x9d, 0x7f})); err != io.EOF {
		t.Errorf("truncated varint: err = %v, want io.EOF", err)
	}
	if _, err := readQUICVarint(bytes.NewReader(nil)); err != io.EOF {
		t.Errorf("empty input: err = %v, want io.EOF", err)
	}
}
