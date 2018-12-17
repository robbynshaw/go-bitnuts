package bits

import (
	"fmt"
	"strings"
	"unicode"
)

// CreateByteFromString takes a string of 1's and 0's
// and creates a byte. Must not be more than 8 characters
func CreateByteFromString(bt string) byte {
	if len(bt) > 8 {
		panic("Byte cannot be more than 8 characters")
	}

	retVal := byte(0)
	i := 0
	for _, char := range bt {
		s := string(char)
		if s == "1" {
			retVal |= byte(1 << uint(7-i))
		}
		i++
	}
	return retVal
}

// CreateBytesFromString takes a string of 1's and 0's
// and creates a slice of bytes.
func CreateBytesFromString(bts string) []byte {
	ar := make([]byte, 0)
	bts = stripWhiteSpace(bts)
	for len(bts) > 0 {
		btstr := bts[0:8]
		bts = bts[8:]
		bt := CreateByteFromString(btstr)
		ar = append(ar, bt)
	}
	return ar
}

// SprintBytes returns the string representation of a byte slice
func SprintBytes(bts []byte) string {
	s := ""
	for _, bt := range bts {
		s += fmt.Sprintf("%08b ", bt)
	}
	return s[0 : len(s)-1]
}

// From https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
func stripWhiteSpace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
