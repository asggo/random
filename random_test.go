package random

import (
	"math"
	"strings"
	"testing"
)

type char struct {
	charset string
	length  uint64
}

type byteint struct {
	bytearray []byte
	result    uint64
}

type intrange struct {
	start uint64
	end   uint64
}

func TestUint64Range(t *testing.T) {
	var tests = []intrange{
		{0, 10},
		{111, 112},
		{0, 1000000},
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

	_, err := Chars("", 15)
	if err == nil {
		t.Error("An empty character set should produce an error.")
	}

	_, err = Chars("ABCDEF", 0)
	if err == nil {
		t.Error("N = 0 should produce an error.")
	}

	for _, char := range chars {
		s, _ := Chars(char.charset, char.length)
		if uint64(len(s)) != char.length {
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
		n, err := Uint8()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < 0 || n > math.MaxUint8 {
			t.Error("Expected integer within range of", 0, "and", math.MaxUint8, "got", n)
		}
	}
}

func TestInt8(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Int8()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < math.MinInt8 || n > math.MaxInt8 {
			t.Error("Expected integer within range of", math.MinInt8, "and", math.MaxInt8, "got", n)
		}
	}
}

func TestUint16(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Uint16()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < 0 || n > math.MaxUint16 {
			t.Error("Expected integer within range of", 0, "and", math.MaxUint16, "got", n)
		}
	}
}

func TestInt16(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Int16()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < math.MinInt16 || n > math.MaxInt16 {
			t.Error("Expected integer within range of", math.MinInt16, "and", math.MaxInt16, "got", n)
		}
	}
}

func TestUint32(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Uint32()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < 0 || n > math.MaxUint32 {
			t.Error("Expected integer within range of", 0, "and", math.MaxUint32, "got", n)
		}
	}
}

func TestInt32(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Int32()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < math.MinInt32 || n > math.MaxInt32 {
			t.Error("Expected integer within range of", math.MinInt32, "and", math.MaxInt32, "got", n)
		}
	}
}

func TestUint64(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Uint64()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < 0 || n > math.MaxUint64 {
			t.Error("Expected integer within range of", 0, "and 18446744073709551615 got", n)
		}
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < 100; i++ {
		n, err := Int64()
		if err != nil {
			t.Error("Unexpected Error:", err)
		}

		if n < math.MinInt64 || n > math.MaxInt64 {
			t.Error("Expected integer within range of", math.MinInt64, "and", math.MaxInt64, "got", n)
		}
	}
}
