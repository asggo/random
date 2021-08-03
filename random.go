// The random package provides cryptographically secure random integers
// and strings.
package random

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"math"
)

// Uint64Range returns a random 64-bit unsigned integer in the range [start, end].
// An error is returned if start is greater than end.
func Uint64Range(start, end uint64) (uint64, error) {
	var val uint64
	var err error

	if start >= end {
		return val, errors.New("Start value must be less than end value.")
	}

	// Get uniformly distributed numbers in the range 0 to size. Using the
	// arc4random_uniform algorithm.
	// https://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/lib/libc/crypt/arc4random_uniform.c
	size := end - start // Get range size
	min := (math.MaxUint64 - size) % size

	for {
		val, err = Uint64()
		if err != nil {
			return val, err
		}

		if val > min {
			break
		}
	}

	val = val % size
	// End arc4random_uniform

	// Add start to val to shift numbers to correct range.
	return val + start, nil
}

// Chars returns a random string of length n, which consists of the given
// character set. If the charset is empty or n is less than or equal to zero
// then an empty string is returned.
func Chars(charset string, n uint64) (string, error) {
	if n == 0 {
		return "", errors.New("Requested string length cannot be 0.")
	}

	if len(charset) == 0 {
		return "", errors.New("Charset cannot be empty.")
	}

	length := uint64(len(charset))
	b := make([]byte, n)

	for i := range b {
		j, err := Uint64Range(0, length)
		if err != nil {
			return "", err
		}
		b[i] = charset[j]
	}

	return string(b), nil
}

// Alpha returns a string of length n, which consists of random upper case and
// lowercase characters. If n is less than or equal to zero then an empty
// string is returned
func Alpha(n uint64) (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return Chars(charset, n)
}

// AlphaNum returns a string of length n, which consists of random uppercase,
// lowercase, and numeric characters. If n is zero then an empty string is
// returned.
func AlphaNum(n uint64) (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	return Chars(charset, n)
}

// Token returns a string suitable for cryptographic tokens such as session ids.
func Token() (string, error) {
	var bytes [32]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes[:]), nil
}

// Uint8 returns a random 8-bit unsigned integer. Return 0 and an error if
// unable to get random data.
func Uint8() (uint8, error) {
	var bytes [1]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint8(0), err
	}

	return bytes[0], nil
}

// Int8 returns a random 8-bit signed integer. Return 0 and an error if
// unable to get random data.
func Int8() (int8, error) {
	i, err := Uint8()

	if err != nil {
		return int8(0), err
	}

	return int8(i), nil
}

// Uint16 returns a random 16-bit unsigned integer. Return 0 and an error if
// unable to get random data.
func Uint16() (uint16, error) {
	var bytes [2]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint16(0), err
	}

	return binary.LittleEndian.Uint16(bytes[:]), nil
}

// Int16 returns a random 16-bit signed integer. Return 0 and an error if
// unable to get random data.
func Int16() (int16, error) {
	i, err := Uint16()

	if err != nil {
		return int16(0), err
	}

	return int16(i), nil
}

// Uint32 returns a random 32-bit unsigned integer. Return 0 and an error if
// unable to get random data.
func Uint32() (uint32, error) {
	var bytes [4]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint32(0), err
	}

	return binary.LittleEndian.Uint32(bytes[:]), nil
}

// Int32 returns a random 32-bit signed integer. Return 0 and an error if
// unable to get random data.
func Int32() (int32, error) {
	i, err := Uint32()

	if err != nil {
		return int32(0), err
	}

	return int32(i), nil
}

// Uint64 returns a random 64-bit unsigned integer. Return 0 and an error if
// unable to get random data.
func Uint64() (uint64, error) {
	var bytes [8]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint64(0), err
	}

	return binary.LittleEndian.Uint64(bytes[:]), nil
}

// Int64 returns a random 64-bit signed integer. Return 0 and an error if
// unable to get random data.
func Int64() (int64, error) {
	i, err := Uint64()

	if err != nil {
		return int64(0), err
	}

	return int64(i), nil
}
