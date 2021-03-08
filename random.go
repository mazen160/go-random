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
// Provided max can't be <= 0.
func GetInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("can't define input as <=0")
	}
	nbig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return max, err
	}
	n := int(nbig.Int64())

	return n, err
}

// GetIntInsecure generate a random integer using a seed of current system time.
func GetIntInsecure(i int) int {
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

// StringRange generates a secure random string within the given range.
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
			mrange = GetIntInsecure(len(charset))
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
func Choice(j []string) (string, error) {
	i, err := GetInt(len(j))
	if err != nil {
		return "", err
	}

	return j[i], nil
}

// Choice makes a random choice from a slice of string.
// Use only when the random choice does not require to be secure.
func ChoiceInsecure(j []string) string {
	i := GetIntInsecure(len(j))

	return j[i]
}
