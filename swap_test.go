package bits

import "testing"

func TestSwapHrzNonOverlapSingleRowSimple(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000001 11110000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00001111 10000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
	}
	act := SwapHorizontal(org, 7, 1, 5, 1)

	assertMatrixMatches(t, exp, act, 7, true)
}

func TestSwapHrzNonOverlapSingleRow(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000001 11110010 10000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 01011111 10000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
	}
	act := SwapHorizontal(org, 7, 1, 5, 1)

	assertMatrixMatches(t, exp, act, 7, true)
}

func TestSwapHrzNonOverlapMultiRow(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000101 10001000 00000000"),
		CreateBytesFromString("00000000 00000011 00111000 00000000"),
		CreateBytesFromString("00000000 00000110 01101000 00000000"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 11011000 00000000"),
		CreateBytesFromString("00000000 00000011 10110000 00000000"),
		CreateBytesFromString("00000000 00000110 11100000 00000000"),
	}
	act := SwapHorizontal(org, 13, 1, 4, 3)

	assertMatrixMatches(t, exp, act, 20, true)
}

func TestSwapHrzNonOverlapMultiRowOverheight(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000101 10001000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000011 00111000 00000000"),
		CreateBytesFromString("00000000 00000110 01101000 00000000"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 11011000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000011 10110000 00000000"),
		CreateBytesFromString("00000000 00000110 11100000 00000000"),
	}
	act := SwapHorizontal(org, 13, 2, 4, 3)

	assertMatrixMatches(t, exp, act, 20, true)
}

func TestSwapHrzOverlapSingleRowSimple(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00111110"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("11110000 00000000 00000000 00000001"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
	}
	act := SwapHorizontal(org, 26, 1, 5, 1)

	assertMatrixMatches(t, exp, act, 7, true)
}

func TestSwapHrzOverlapSingleRow(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00100000 00000000 00000101 11011001"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("10100000 00000000 00000100 10011011"),
	}
	act := SwapHorizontal(org, 21, 3, 7, 1)

	assertMatrixMatches(t, exp, act, 21, true)
}

func TestSwapHrzOverlapMultiRow(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("01000000 00000000 00000001 01100001"),
		CreateBytesFromString("00000000 00000000 00000111 11110000"),
		CreateBytesFromString("00100000 00000000 00000101 11011001"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("11000000 00000000 00000000 10100010"),
		CreateBytesFromString("11100000 00000000 00000000 00001111"),
		CreateBytesFromString("10100000 00000000 00000100 10011011"),
	}
	act := SwapHorizontal(org, 21, 1, 7, 3)

	assertMatrixMatches(t, exp, act, 21, true)
}

func TestSwapHrzLeftOverlapMultiRow(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("10000000 00000000 00000000 00000111"),
		CreateBytesFromString("01111000 00000000 00000000 00000000"),
		CreateBytesFromString("01001000 00000000 00000000 00000101"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("01111000 00000000 00000000 00000000"),
		CreateBytesFromString("10000000 00000000 00000000 00000111"),
		CreateBytesFromString("11010000 00000000 00000000 00000100"),
	}
	act := SwapHorizontal(org, 29, 1, 4, 3)

	assertMatrixMatches(t, exp, act, 29, true)
}

func TestSwapHrzOverlapMultiRowOverheight(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("01000000 00000000 00000001 01100001"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("00000000 00000000 00000111 11110000"),
		CreateBytesFromString("00100000 00000000 00000101 11011001"),
	}
	exp := [][]byte{
		CreateBytesFromString("11000000 00000000 00000000 10100010"),
		CreateBytesFromString("00000000 00000000 00000000 00000000"),
		CreateBytesFromString("11100000 00000000 00000000 00001111"),
		CreateBytesFromString("10100000 00000000 00000100 10011011"),
	}
	act := SwapHorizontal(org, 21, 2, 7, 3)

	assertMatrixMatches(t, exp, act, 21, true)
}

func TestSwapVrtNonOverlapMultiRow(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00000000 00100000 01000000 00000000"),
		CreateBytesFromString("00100000 01111100 01010101 00000010"),
		CreateBytesFromString("00000000 10101011 00111000 10000000"),
		CreateBytesFromString("00000000 00100100 00001000 10000000"),
	}
	exp := [][]byte{
		CreateBytesFromString("00000000 00101011 00111000 10000000"),
		CreateBytesFromString("00100000 00100100 00001000 10000010"),
		CreateBytesFromString("00000000 10100000 01000000 00000000"),
		CreateBytesFromString("00000000 01111100 01010101 00000000"),
	}
	act := SwapVertical(org, 9, 1, 16, 2)

	assertMatrixMatches(t, exp, act, 9, true)
}

func TestSwapVrtOverlapMultiRow(t *testing.T) {
	org := [][]byte{
		CreateBytesFromString("00100000 00000000 00100000 01000000"),
		CreateBytesFromString("11010000 00100000 01111100 01010101"),
		CreateBytesFromString("01100000 00000000 10101011 00111000"),
		CreateBytesFromString("10010000 00000000 00100100 00001000"),
	}
	exp := [][]byte{
		CreateBytesFromString("01100000 00000000 00101011 00111000"),
		CreateBytesFromString("10010000 00100000 00100100 00001000"),
		CreateBytesFromString("00100000 00000000 10100000 01000000"),
		CreateBytesFromString("11010000 00000000 01111100 01010101"),
	}
	act := SwapVertical(org, 17, 1, 19, 2)

	assertMatrixMatches(t, exp, act, 17, true)
}

func TestSwapVrtNoOutOfRange(t *testing.T) {
	matrix := make([][]byte, 24)
	row := CreateBytesFromString("00100000 00000000 00100000")
	for i := 0; i < 24; i++ {
		matrix[i] = row
	}
	SwapVertical(matrix, 0, 14, 1, 11)
}
