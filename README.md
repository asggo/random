Random
======

package random
    import "github.com/averagesecurityguy/random"

    The random package provides cryptographically secure random integers
    and strings

Functions
=========

func Alpha(n uint64) (string, error)

    Alpha returns a string of length n, which consists of random upper case
    and lowercase characters. If n is less than or equal to zero then an
    empty string is returned

func AlphaNum(n uint64) (string, error)

    AlphaNum returns a string of length n, which consists of random
    uppercase, lowercase, and numeric characters. If n is zero then an empty
    string is returned.

func Chars(charset string, n uint64) (string, error)

    Chars returns a random string of length n, which consists of the given
    character set. If the charset is empty or n is less than or equal to
    zero then an empty string is returned.

func Int16() (int16, error)

    Int16 returns a random 16-bit signed integer. Return 0 and an error if
    unable to get random data.

func Int32() (int32, error)

    Int32 returns a random 32-bit signed integer. Return 0 and an error if
    unable to get random data.

func Int64() (int64, error)

    Int64 returns a random 64-bit signed integer. Return 0 and an error if
    unable to get random data.

func Int8() (int8, error)

    Int8 returns a random 8-bit signed integer. Return 0 and an error if
    unable to get random data.

func Token() (string, error)

    Token returns a string suitable for cryptographic tokens such as session
    ids.

func Uint16() (uint16, error)

    Uint16 returns a random 16-bit unsigned integer. Return 0 and an error
    if unable to get random data.

func Uint32() (uint32, error)

    Uint32 returns a random 32-bit unsigned integer. Return 0 and an error
    if unable to get random data.

func Uint64() (uint64, error)

    Uint64 returns a random 64-bit unsigned integer. Return 0 and an error
    if unable to get random data.

func Uint64Range(start, end uint64) (uint64, error)

    Uint64Range returns a random 64-bit unsigned integer in the range
    [start, end]. An error is returned if start is greater than end.

func Uint8() (uint8, error)

    Uint8 returns a random 8-bit unsigned integer. Return 0 and an error if
    unable to get random data.
