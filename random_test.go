package random

import (
	"math"
	"strings"
	"testing"
)

type char struct {
	charset string
	length  int
}

type byteint struct {
	bytearray []byte
	result    uint64
}

type intrange struct {
	start uint64
	end   uint64
}

func TestRandomBytes(t *testing.T) {
	var tests = []int{1, 5, 10}

	for _, test := range tests {
		r := Bytes(test)
		if len(r) != test {
			t.Error("Expected", test, "got", len(r))
		}
	}
}

func TestbytesToInt(t *testing.T) {
	var tests = []byteint{
		{[]byte{128, 0, 0, 0, 0, 0, 0, 0}, 9223372036854775808},
		{[]byte{0, 0, 0, 0, 128, 0, 0, 0}, 2147483648},
		{[]byte{0, 0, 0, 0, 0, 0, 0, 1}, 1},
		{[]byte{0, 0, 0, 0, 0, 3}, 3},
		{[]byte{0, 0, 0, 5}, 5},
	}

	for _, test := range tests {
		i := bytesToInt(test.bytearray)
		if i != test.result {
			t.Error("Expected", test.result, "got", i)
		}
	}
}

func TestUint64Range(t *testing.T) {
	var tests = []intrange{
		{0, 10},
		{9223372036854775800, 9223372036854775808},
	}

	_, err := Uint64Range(10, 0)
	if err == nil {
		t.Error("Start greater than end should produce an error.")
	}

	for _, test := range tests {
		i, _ := Uint64Range(test.start, test.end)
		if i < test.start || i > test.end {
			t.Error("Expected integer within range of", test.start, "and", test.end, "got", i)
		}
	}
}

func TestChars(t *testing.T) {
	var chars = []char{
		{"abcdefghijklmnopqrstuvwxyz", 5},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 20},
		{"0123456789abcdef", 15},
	}

	s := Chars("", 15)
	if s != "" {
		t.Error("An empty character set should produce an empty string.")
	}

	s = Chars("ABCDEF", 0)
	if s != "" {
		t.Error("N <= 0 should produce an empty string.")
	}

	for _, char := range chars {
		s := Chars(char.charset, char.length)
		if len(s) != char.length {
			t.Error("Expected string of length", char.length, "and got string of length", len(s))
		}

		for _, c := range s {
			if !strings.Contains(char.charset, string(c)) {
				t.Error("Expected a character in", char.charset, "got", c)
			}
		}
	}
}

func TestUint8(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Uint8()
		if n < 0 || n > math.MaxUint8 {
			t.Error("Expected integer within range of", 0, "and", math.MaxUint8-1, "got", n)
		}
	}
}

func TestInt8(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Int8()
		if n < math.MinInt8 || n > math.MaxInt8 {
			t.Error("Expected integer within range of", math.MinInt8, "and", math.MaxInt8, "got", n)
		}
	}
}

func TestUint16(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Uint16()
		if n < 0 || n > math.MaxUint16 {
			t.Error("Expected integer within range of", 0, "and", math.MaxUint16, "got", n)
		}
	}
}

func TestInt16(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Int16()
		if n < math.MinInt16 || n > math.MaxInt16 {
			t.Error("Expected integer within range of", math.MinInt16, "and", math.MaxInt16, "got", n)
		}
	}
}

func TestUint32(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Uint32()
		if n < 0 || n > math.MaxUint32 {
			t.Error("Expected integer within range of", 0, "and", math.MaxUint32, "got", n)
		}
	}
}

func TestInt32(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Int32()
		if n < math.MinInt32 || n > math.MaxInt32 {
			t.Error("Expected integer within range of", math.MinInt32, "and", math.MaxInt32, "got", n)
		}
	}
}

func TestUint64(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Uint64()
		if n < 0 || n > math.MaxUint64 {
			t.Error("Expected integer within range of", 0, "and 18446744073709551615 got", n)
		}
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := Int64()
		if n < math.MinInt64 || n > math.MaxInt64 {
			t.Error("Expected integer within range of", math.MinInt64, "and", math.MaxInt64, "got", n)
		}
	}
}

func TestUint64n(t *testing.T) {
	for i := 0; i < 100; i++ {
		max := uint64(i * i)
		n := Uint64n(max)
		if n < 0 || n > uint64(max) {
			t.Error("Expected integer within range of", 0, "and", max, "got", n)
		}
	}
}
