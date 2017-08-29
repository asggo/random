Random
======

package random
    import "github.com/averagesecurityguy/random"

    The random package provides cryptographically secure random integers
    and strings

Functions
=========

func Bytes(size int) []byte

    Bytes returns an array of the specified size filled with random bytes.

func Chars(charset string, n int) string

    Chars returns a random string using the given character set

func Int16() int16

    Int16 returns a random 16-bit signed integer.

func Int32() int32

    Int32 returns a random 32-bit signed integer.

func Int64() int64

    Int64 returns a random 64-bit signed integer.

func Int64n(n uint64) int64

    Int64n returns a random 64-bit signed integer in the range [0, n)

func Int8() int8

    Int8 returns a random 8-bit signed integer.

func Uint16() uint16

    Uint16 returns a random 16-bit unsigned integer.

func Uint32() uint32

    Uint32 returns a random 32-bit unsigned integer.

func Uint64() uint64

    Uint64 returns a random 64-bit unsigned integer.

func Uint64Range(start, end uint64) (uint64, error)

    Uint64Range returns a random 64-bit unsigned integer in the range
    [start, end)

func Uint64n(n uint64) uint64

    Uint64n returns a random 64-bit unsigned integer in the range [0, n)

func Uint8() uint8

    Uint8 returns a random 8-bit unsigned integer.
