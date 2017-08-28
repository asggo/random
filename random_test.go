package cryptanalysis

import (
	"testing"
)

type byteint struct{
	bytearray []byte
	result    uint64
}

type intrange struct{
	start uint64
	end   uint64
}

func TestRandomBytes(t *testing.T) {
	var tests = []int{1, 5, 10,}

	for _, test := range tests {
		r, _ := RandomBytes(test)
		if len(r) != test {
			t.Error("Expected", test, "got", len(r))
		}
	}
}


func TestBytesToInt(t *testing.T) {
	var tests = []byteint{
		{[]byte{128,0,0,0,0,0,0,0}, 9223372036854775808},
		{[]byte{0,0,0,0,128,0,0,0}, 2147483648},
		{[]byte{0,0,0,0,0,0,0,1}, 1},
		{[]byte{0,0,0,0,0,3}, 3},
		{[]byte{0,0,0,5}, 5},


	}

	_, err := BytesToInt([]byte{})
	if err == nil {
		t.Error("Empty byte array should produce an error.")
	}

	_, err = BytesToInt([]byte{0,0,0,0,0,0,0,0,0})
	if err == nil {
		t.Error("Long byte array should produce an error.")
	}

	for _, test := range tests {
		i, _ := BytesToInt(test.bytearray)
		if i != test.result {
			t.Error("Expected", test.result, "got", i)
		}
	}
}


func TestRandomIntRange(t *testing.T) {
	var tests = []intrange{
		{0, 10},
		{9223372036854775800, 9223372036854775808},
	}

	_, err := RandomIntRange(10, 0)
	if err == nil {
		t.Error("Start greater than end should produce an error.")
	}

	for _, test := range tests {
		i, _:= RandomIntRange(test.start, test.end)
		if i < test.start || i > test.end {
			t.Error("Expected integer within range of", test.start, "and", test.end, "got", i)
		}
	}
}
