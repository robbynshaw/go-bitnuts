package bits

import (
	"testing"
)

func TestUnset(t *testing.T) {
	a := byte(255)
	b := Unset(a, 3)
	c := CreateByteFromString("11101111")

	assertBytesMatch(t, b, c)
}

func TestSet(t *testing.T) {
	a := byte(0)
	b := Set(a, 3)
	c := CreateByteFromString("00010000")

	assertBytesMatch(t, b, c)
}

func TestSetIsUnchanged(t *testing.T) {
	a := CreateByteFromString("00010000")
	b := Set(a, 3)
	c := CreateByteFromString("00010000")

	assertBytesMatch(t, b, c)
}

func TestIsSetInRowTrueAtBeginning(t *testing.T) {
	a := CreateBytesFromString("00000000 00000000 10000000 00000000")
	b := IsSetInRow(a, 16)

	assertBool(t, true, b)
}

func TestIsSetInRowTrueAtEnd(t *testing.T) {
	a := CreateBytesFromString("00000000 00000001 00000000 00000000")
	b := IsSetInRow(a, 15)

	assertBool(t, true, b)
}

func TestIsSetInRowTrueInMiddle(t *testing.T) {
	a := CreateBytesFromString("00000000 00001000 00000000 00000000")
	b := IsSetInRow(a, 12)

	assertBool(t, true, b)
}

func TestIsUnSetInRowTrueAtBeginning(t *testing.T) {
	a := CreateBytesFromString("11111111 11111111 01111111 11111111")
	b := IsSetInRow(a, 16)

	assertBool(t, false, b)
}

func TestIsUnSetInRowTrueAtEnd(t *testing.T) {
	a := CreateBytesFromString("11111111 11111110 11111111 11111111")
	b := IsSetInRow(a, 15)

	assertBool(t, false, b)
}

func TestIsUnSetInRowTrueInMiddle(t *testing.T) {
	a := CreateBytesFromString("11111111 11110111 11111111 11111111")
	b := IsSetInRow(a, 12)

	assertBool(t, false, b)
}

func TestDoSetInRowTrueAtBeginning(t *testing.T) {
	org := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00000000 00000000 10000000 00000000")
	act := SetInRow(org, 16)

	assertByteSliceMatches(t, exp, act)
}

func TestDoSetInRowTrueAtEnd(t *testing.T) {
	org := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00000000 00000001 00000000 00000000")
	act := SetInRow(org, 15)

	assertByteSliceMatches(t, exp, act)
}

func TestDoSetInRowTrueInMiddle(t *testing.T) {
	org := CreateBytesFromString("00000000 00000000 00000000 00000000")
	exp := CreateBytesFromString("00000000 00001000 00000000 00000000")
	act := SetInRow(org, 12)

	assertByteSliceMatches(t, exp, act)
}

func TestDoUnSetInRowTrueAtBeginning(t *testing.T) {
	org := CreateBytesFromString("11111111 11111111 11111111 11111111")
	exp := CreateBytesFromString("11111111 11111111 01111111 11111111")
	act := UnsetInRow(org, 16)

	assertByteSliceMatches(t, exp, act)
}

func TestDoUnSetInRowTrueAtEnd(t *testing.T) {
	org := CreateBytesFromString("11111111 11111111 11111111 11111111")
	exp := CreateBytesFromString("11111111 11111110 11111111 11111111")
	act := UnsetInRow(org, 15)

	assertByteSliceMatches(t, exp, act)
}

func TestDoUnSetInRowTrueInMiddle(t *testing.T) {
	org := CreateBytesFromString("11111111 11111111 11111111 11111111")
	exp := CreateBytesFromString("11111111 11110111 11111111 11111111")
	act := UnsetInRow(org, 12)

	assertByteSliceMatches(t, exp, act)
}

func assertBool(t *testing.T, expected, actual bool) {
	if expected != actual {
		t.Errorf("Expected %v, received %v\n", expected, actual)
	}
}

func assertBytesMatch(t *testing.T, a, b byte) {
	if a != b {
		t.Errorf("Bytes should match: %08b - %08b\n", a, b)
	}
}
