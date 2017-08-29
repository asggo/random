// The random package provides cryptographically secure random integers
// and strings.
package random

import (
	"crypto/rand"
	"errors"
	"math"
)

// Bytes returns an array of the specified size filled with random bytes.
func Bytes(size int) []byte {
	array := make([]byte, size)
	_, err := rand.Read(array)

	if err != nil {
		panic(err)
	}

	return array
}

// BytesToInt returns a uint64 created from an array of 1 to 8 bytes.
func bytesToInt(array []byte) uint64 {
	var integer uint64
	size := len(array)

	for i, b := range array {
		shift := uint64((size - i - 1) * 8)
		integer |= uint64(b) << shift
	}

	return integer
}

// Uint64Range returns a random 64-bit unsigned integer in the range [start, end)
func Uint64Range(start, end uint64) (uint64, error) {
	val := uint64(0)

	if start >= end {
		return val, errors.New("Start value must be less than end value.")
	}

	for {
		val = bytesToInt(Bytes(8))
		if val <= (math.MaxUint64 - (math.MaxUint64 % end)) {
			break
		}
	}

	val = val % (end - start)
	val = val + start

	return val, nil
}

// Chars returns a random string using the given character set
func Chars(charset string, n int) string {
	if n == 0 {
		return ""
	}

	if len(charset) == 0 {
		return ""
	}

	length := uint64(len(charset))
	b := make([]byte, n)

	for i := range b {
		b[i] = charset[Uint64n(length)]
	}

	return string(b)
}

// Uint8 returns a random 8-bit unsigned integer.
func Uint8() uint8 {
	i, _ := Uint64Range(0, math.MaxUint8)

	return uint8(i)
}

// Int8 returns a random 8-bit signed integer.
func Int8() int8 {
	i, _ := Uint64Range(0, math.MaxUint8)

	return int8(i)
}

// Uint16 returns a random 16-bit unsigned integer.
func Uint16() uint16 {
	i, _ := Uint64Range(0, math.MaxUint16)

	return uint16(i)
}

// Int16 returns a random 16-bit signed integer.
func Int16() int16 {
	i, _ := Uint64Range(0, math.MaxUint16)

	return int16(i)
}

// Uint32 returns a random 32-bit unsigned integer.
func Uint32() uint32 {
	i, _ := Uint64Range(0, math.MaxUint32)

	return uint32(i)
}

// Int32 returns a random 32-bit signed integer.
func Int32() int32 {
	i, _ := Uint64Range(0, math.MaxUint32)

	return int32(i)
}

// Uint64 returns a random 64-bit unsigned integer.
func Uint64() uint64 {
	i, _ := Uint64Range(0, math.MaxUint64)

	return i
}

// Int64 returns a random 64-bit signed integer.
func Int64() int64 {
	i, _ := Uint64Range(0, math.MaxUint64)

	return int64(i)
}

// Uint64n returns a random 64-bit unsigned integer in the range [0, n)
func Uint64n(n uint64) uint64 {
	i, _ := Uint64Range(0, n)

	return i
}

// Int64n returns a random 64-bit signed integer in the range [0, n)
func Int64n(n uint64) int64 {
	i, _ := Uint64Range(0, n)

	for {
		v := int64(i)

		if v >= 0 {
			return v
		}

		i, _ = Uint64Range(0, n)
	}
}
