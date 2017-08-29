Random
======

package random
    import "github.com/averagesecurityguy/random"

    The random package provides cryptographically secure random integers
    and strings

Functions
=========

func Alpha(n int) string

    Alpha returns a string of length n, which consists of random upper case
    and lowercase characters. If n is less than or equal to zero then an
    empty string is returned

func AlphaNum(n int) string

    AlphaNum returns a string of length n, which consists of random
    uppercase, lowercase, and numeric characters. If n is zero then an empty
    string is returned.

func Bytes(size int) []byte

    Bytes returns an array of the specified size filled with random bytes.
    Bytes will panic if random bytes cannot be read from the OS.

func Chars(charset string, n int) string

    Chars returns a random string of length n, which consists of the given
    character set. If the charset is empty or n is less than or equal to
    zero then an empty string is returned.

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
    [start, end). An error is returned if start is greater than end.

func Uint64n(n uint64) uint64

    Uint64n returns a random 64-bit unsigned integer in the range [0, n)

func Uint8() uint8

    Uint8 returns a random 8-bit unsigned integer.