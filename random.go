package random

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var (
	Digits                string = "0123456789"                                                         // Digits: [0-9]
	ASCIILettersLowercase string = "abcdefghijklmnopqrstuvwxyz"                                         // Asci Lowerrcase Letters: [a-z]
	ASCIILettersUppercase string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"                                         // Ascii Uppercase Letters: [A-Z]
	Letters               string = ASCIILettersLowercase + ASCIILettersUppercase                        // Ascii Letters: [a-zA-Z]
	ASCIICharacters       string = ASCIILettersLowercase + ASCIILettersUppercase + Digits               // Ascii Charaters: [a-zA-Z0-9]
	Hexdigits             string = "0123456789abcdefABCDEF"                                             // Hex Digits: [0-9a-fA-F]
	Octdigits             string = "01234567"                                                           // Octal Digits: [0-7]
	Punctuation           string = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                                 // Punctuation and special characters
	Printables            string = Digits + ASCIILettersLowercase + ASCIILettersUppercase + Punctuation // Printables
)

// GetInt generates a cryptographically-secure random Int.
// If fails, it returns the max int and an error.
func GetInt(i int) (int, error) {
	var n int
	defer func() {
		if err := recover(); err != nil {
			n = i
		}
	}()
	nbig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(i)))
	n = int(nbig.Int64())

	if n <= 0 {
		return GetInt(i)
	}

	return n, err
}

// GetMathInt generate a random integer using a seed of current system time.
func GetMathInt(i int) int {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	return seededRand.Intn(i)
}

// String generates a cryptographically secure string.
func String(n int) (string, error) {
	return Random(n, ASCIICharacters, true)
}

// String generates a cryptographically insecure string.
// Use only when generating random data that does not require to be secure.
func StringInsecure(n int) (string, error) {
	return Random(n, ASCIICharacters, false)
}

// StringRange generates a secure random string with the given range.
func StringRange(min int, max int) (string, error) {
	i, err := IntRange(min, max)
	if err != nil {
		return "", err
	}
	return String(i)
}

// IntRange returns a random integer between a given range.
func IntRange(min int, max int) (int, error) {
	i, err := GetInt(max - min)

	if err != nil {
		return max, fmt.Errorf("error getting safe int with crypto/rand")
	}
	i += min
	return i, nil
}

// Random is responsible for generating Random data from a given character set.
func Random(n int, charset string, isSecure bool) (string, error) {
	var charsetByte = []byte(charset)

	s := make([]byte, n)

	var mrange int
	var err error
	for i := range s {
		if isSecure {
			mrange, err = GetInt(len(charset))
			if err != nil {
				return "", fmt.Errorf("error getting safe int with crypto/rand")
			}
		} else {
			mrange = GetMathInt(len(charset))
		}

		s[i] = charsetByte[mrange]
	}

	return string(s), nil
}

// Bytes generates a cryptographically secure set of bytes.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crypto_rand.Read(b)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}

// Choice makes a random choice from a slice of string.
func Choice(j []string) string {
	i := GetMathInt(len(j))

	return j[i]
}
