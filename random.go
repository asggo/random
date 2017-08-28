// The random package provides cryptographically secure random integers,
// strings, and 
package random

import (
    "errors"
    "math"
    "crypto/rand"
)


// RandomBytes returns an array of the specified size filled with random bytes.
func RandomBytes(size int) ([]byte, error) {
    array := make([]byte, size)
    _, err := rand.Read(array)

    if err != nil {
        return array, err
    }

	return array, nil
}


// BytesToInt returns a uint64 created from an array of 1 to 8 bytes.
func bytesToInt(array []byte) (uint64, error) {
	if len(array) == 0 || len(array) > 8 {
        return uint64(0), errors.New("Need 1 to 8 bytes to create integer.")
    }

    var integer uint64
    size := len(array)

    for i, b := range array {
		shift := uint64((size - i - 1) * 8)
		integer |= uint64(b) << shift
	}

    return integer, nil
}


// Uint8 returns a random 8-bit unsigned integer.
func Uint8() uint8 {
    b, err := RandomBytes(1)
    if err != nil {
        return val, err
    }

    return uint8(b[0])
}

// Uint16 returns a random 16-bit unsigned integer.
func Uint16() uint16 {
    var val uint64

    b, err := RandomBytes(2)
    if err != nil {
        return val, err
    }

    val, err = bytesToInt(b)
    if err != nil {
        return val, err
    }

    return uint16(val)
}

// Uint32 returns a random 32-bit unsigned integer.
func Uint32() uint32 {
    b, err := RandomBytes(4)
    if err != nil {
        return val, errors.New("Could not read random bytes from OS.")
    }

    val, err = bytesToInt(b)
    if err != nil {
        return val, errors.New("Could not convert bytes to unsigned integer.")
    }

    return uint32(val)
}

// Uint64 returns a random 64-bit unsigned integer.
func Uint64() uint64 {
    b, err := RandomBytes(2)
    if err != nil {
        return val, errors.New("Could not read random bytes from OS.")
    }

    val, err = bytesToInt(b)
    if err != nil {
        return val, errors.New("Could not convert bytes to unsigned integer.")
    }

    return val
}

// Return a random uint64 in the range [start, end)
func Uint64Range(start, end uint64) (uint64, error) {
	val = uint64(0)

    if start >= end {
        return val, errors.New("Start value must be less than end value.")
    }

    for {
		b, err := RandomBytes(8)
		if err != nil {
			return val, errors.New("Could not read random bytes from OS.")
		}

        val, err = bytesToInt(b)
		if err != nil {
			return val, errors.New("Could not convert bytes to unsigned integer.")
		}

        if (val <= (math.MaxUint64 - (math.MaxUint64 % end))){ break }
    }

    val = val % (end - start)
    val = val + start

    return val, nil
}
