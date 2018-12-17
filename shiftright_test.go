package bits

import (
	"fmt"
	"testing"
)

func printByte(bt byte) {
	fmt.Printf("%08b\n", bt)
}

func printBytes(bts []byte) {
	for _, bt := range bts {
		fmt.Printf("%08b ", bt)
	}
	fmt.Println()
}

func TestEvenRightShiftedBitsDontMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("011010100101110010001101")
	c := ShiftRight(a, 8)

	if byteArraysMatch(b, c) {
		fmt.Println("Byte arrays should not match")
		t.Fail()
	}
}

func TestEvenRightShiftedBitsMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("100011010110101001011100")
	c := ShiftRight(a, 8)

	if !byteArraysMatch(b, c) {
		fmt.Println("Original:")
		printBytes(a)
		fmt.Println("Byte arrays should match:")
		printBytes(b)
		printBytes(c)
		t.Fail()
	}
}

func TestUnevenRightShiftedBitsDontMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("011010100101110010001101")
	c := ShiftRight(a, 5)

	if byteArraysMatch(b, c) {
		fmt.Println("Byte arrays should not match")
		t.Fail()
	}
}

func TestUnevenRightShiftedBitsMatch(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")
	b := CreateBytesFromString("011010110101001011100100")
	c := ShiftRight(a, 5)

	if !byteArraysMatch(b, c) {
		fmt.Println("Original:")
		printBytes(a)
		fmt.Println("Byte arrays should match:")
		printBytes(b)
		printBytes(c)
		t.Fail()
	}
}

func TestZeroRightShiftedPanics(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")

	assertPanic(t, "Zero shift should fail", func() {
		_ = ShiftRight(a, 0)
	})
}

func TestExcessRightShiftedPanics(t *testing.T) {
	a := CreateBytesFromString("011010100101110010001101")

	assertPanic(t, "Excess shift should fail", func() {
		_ = ShiftRight(a, 30)
	})
}
