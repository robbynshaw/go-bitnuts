package bits

import (
	"fmt"
	"testing"
)

func byteArraysMatch(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestEvenShiftedBitsDontMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("011010100101110010001101")
	c := ShiftLeft(a, 8)

	if byteArraysMatch(b, c) {
		fmt.Println("Byte arrays should not match")
		t.Fail()
	}
}

func TestEvenShiftedBitsMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("010111001000110101101010")
	c := ShiftLeft(a, 8)

	if !byteArraysMatch(b, c) {
		fmt.Println("Byte arrays should match")
		t.Fail()
	}
}

func TestUnevenShiftedBitsDontMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("011010100101110010001101")
	c := ShiftLeft(a, 5)

	if byteArraysMatch(b, c) {
		fmt.Println("Byte arrays should not match")
		t.Fail()
	}
}

func TestUnvenShiftedBitsMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("010010111001000110101101")
	c := ShiftLeft(a, 5)

	if !byteArraysMatch(b, c) {
		fmt.Println("Byte arrays should match")
		t.Fail()
	}
}

func TestZeroShiftedPanics(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")

	assertPanic(t, "Zero shift should fail", func() {
		_ = ShiftLeft(a, 0)
	})
}

func TestExcessShiftedPanics(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")

	assertPanic(t, "Excess shift should fail", func() {
		_ = ShiftLeft(a, 30)
	})
}
