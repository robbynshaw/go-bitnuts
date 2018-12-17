package bits

import (
	"fmt"
	"testing"
)

func assertPanic(t *testing.T, msg string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(msg)
		}
	}()
	f()
}

func TestExcessBytesFromStringFails(t *testing.T) {
	assertPanic(t, "Blah", func() {
		_ = CreateByteFromString("01010001010010101")
	})
}

func TestByteFromStringNotMatches(t *testing.T) {
	a := CreateByteFromString("10101011")
	b := byte(13)

	if a == b {
		fmt.Printf("Bytes should not match: %08b - %08b\n", a, b)
		t.Fail()
	}
}

func TestByteFromStringMatches(t *testing.T) {
	a := CreateByteFromString("10101011")
	b := byte(171)

	if a != b {
		fmt.Printf("Bytes should match: %08b - %08b\n", a, b)
		t.Fail()
	}
}

func TestBytesFromStringNotMatches(t *testing.T) {
	a := CreateBytesFromString("1010101100101111")
	b := []byte{byte(171), byte(2)}

	if a[0] == b[0] && a[1] == b[1] && len(a) == len(b) {
		fmt.Println("Byte arrays should not match")
		t.Fail()
	}
}

func TestBytesFromStringMatches(t *testing.T) {
	a := CreateBytesFromString("1010101100101111")
	b := []byte{byte(171), byte(47)}

	if a[0] != b[0] || a[1] != b[1] || len(a) != len(b) {
		fmt.Println("Byte arrays should match")
		t.Fail()
	}
}

func TestBytesFromStringWithSpacesMatches(t *testing.T) {
	a := CreateBytesFromString("10101011 00101111")
	b := []byte{byte(171), byte(47)}

	if a[0] != b[0] || a[1] != b[1] || len(a) != len(b) {
		fmt.Println("Byte arrays should match")
		t.Fail()
	}
}
