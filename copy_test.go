package bits

import (
	"fmt"
	"testing"
)

func TestCopyADoesntMatch(t *testing.T) {
	a := CreateByteFromString("01100111")
	b := CreateByteFromString("00001111")
	c := Copy(a, 3, 7)

	if b == c {
		fmt.Printf("Bytes should not match: %08b - %08b\n", b, c)
		t.Fail()
	}
}

func TestCopyBDoesntMatch(t *testing.T) {
	a := CreateByteFromString("01100111")
	b := CreateByteFromString("00000000")
	c := Copy(a, 3, 7)

	if b == c {
		fmt.Printf("Bytes should not match: %08b - %08b\n", b, c)
		t.Fail()
	}
}

func TestCopyAMatches(t *testing.T) {
	a := CreateByteFromString("01100111")
	b := CreateByteFromString("00000111")
	c := Copy(a, 3, 7)

	if b != c {
		fmt.Printf("Bytes should match: %08b - %08b\n", b, c)
		t.Fail()
	}
}

func TestCopyBMatches(t *testing.T) {
	a := CreateByteFromString("01110111")
	b := CreateByteFromString("00110110")
	c := Copy(a, 2, 6)

	if b != c {
		fmt.Printf("Bytes should match: %08b - %08b\n", b, c)
		t.Fail()
	}
}

func TestPasteRightDoesntMatch(t *testing.T) {
	a := CreateByteFromString("00010110")
	b := byte(255)
	c := PasteRight(a, b, 3)

	if b == c {
		fmt.Printf("Bytes should not match: %08b - %08b\n", b, c)
		t.Fail()
	}
}

func TestPasteRightMatches(t *testing.T) {
	a := CreateByteFromString("00000110")
	x := CreateByteFromString("11100110")
	b := byte(255)
	c := PasteRight(a, b, 3)

	if x != c {
		fmt.Printf("Bytes should match: %08b - %08b\n", x, c)
		t.Fail()
	}
}

func TestCopyPasteRowMatchesByteStart(t *testing.T) {
	org := CreateBytesFromString("00000000 00000000 10100000 00000000")
	tor := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("10100000 00000000 00000000 00000000")
	act := CopyPasteRow(org, tor, 16, 0, 3)

	assertByteSliceMatches(t, exp, act)
}

func TestCopyPasteRowMatchesByteStartOverlap(t *testing.T) {
	org := CreateBytesFromString("00000000 00000000 10111000 00000000")
	tor := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00000101 11000000 00000000 00000000")
	act := CopyPasteRow(org, tor, 16, 5, 5)

	assertByteSliceMatches(t, exp, act)
}

func TestCopyPasteRowMatchesByteEnd(t *testing.T) {
	org := CreateBytesFromString("00000000 00000000 00110101 00000000")
	tor := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00110101 00000000 00000000 00000000")
	act := CopyPasteRow(org, tor, 18, 2, 6)

	assertByteSliceMatches(t, exp, act)
}

func TestCopyPasteRowMatchesByteEndOverlap(t *testing.T) {
	org := CreateBytesFromString("00000000 00000000 00110101 11000000")
	tor := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00110101 11000000 00000000 00000000")
	act := CopyPasteRow(org, tor, 18, 2, 8)

	assertByteSliceMatches(t, exp, act)
}

func TestCopyPasteRowMatchesByteMiddle(t *testing.T) {
	org := CreateBytesFromString("00000000 00111000 00000000 00000000")
	tor := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00000000 00000000 00000000 00000111")
	act := CopyPasteRow(org, tor, 10, 29, 3)

	assertByteSliceMatches(t, exp, act)
}

func TestCopyPasteRowMatchesByteMiddleOverlap(t *testing.T) {
	org := CreateBytesFromString("00000000 00110100 00000000 00000000")
	tor := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00000011 01000000 00000000 00000000")
	act := CopyPasteRow(org, tor, 10, 6, 4)

	assertByteSliceMatches(t, exp, act)
}
