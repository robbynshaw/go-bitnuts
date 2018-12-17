package bits

import "testing"

func TestToggleSingleBitAt5(t *testing.T) {
	a := CreateByteFromString("00101100")
	b := CreateByteFromString("00101000")
	c := Toggle(a, 5)

	assertBytesMatch(t, b, c)
}

func TestToggleSingleBitAt2(t *testing.T) {
	a := CreateByteFromString("00001100")
	b := CreateByteFromString("00101100")
	c := Toggle(a, 2)

	assertBytesMatch(t, b, c)
}

func TestToggleRowSectionMatches3_9(t *testing.T) {
	a := CreateBytesFromString("00101100 00110110 11011111")
	b := CreateBytesFromString("00110011 11110110 11011111")
	c := ToggleRowSection(a, 3, 9)

	assertByteSliceMatches(t, b, c)
}

func TestVerticalToggleMatches(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("01101010 01011100 10001101"),
		CreateBytesFromString("011l1000 11110100 10l01101"),
		CreateBytesFromString("01001011 11011110 10101110"),
		CreateBytesFromString("01101110 01010100 10011101"),
		CreateBytesFromString("01101110 01010110 00001001"),
		CreateBytesFromString("11101110 01011100 10001111"),
	}

	b := [][]byte{
		CreateBytesFromString("01101010 01111100 10001101"),
		CreateBytesFromString("011l1000 11010100 10l01101"),
		CreateBytesFromString("01001011 11111110 10101110"),
		CreateBytesFromString("01101110 01110100 10011101"),
		CreateBytesFromString("01101110 01110110 00001001"),
		CreateBytesFromString("11101110 01111100 10001111"),
	}

	printColumnln(t, "Original: ", a, 10)
	c := ToggleColumn(a, 10)
	assertMatrixMatches(t, b, c, 10, true)
}

func TestVerticalToggleSectionMatches(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("01101010 01011100 10001101"),
		CreateBytesFromString("011l1000 11110100 10l01101"),
		CreateBytesFromString("01001011 11011110 10101110"),
		CreateBytesFromString("01101110 01010100 10011101"),
		CreateBytesFromString("01101110 01010110 00001001"),
		CreateBytesFromString("11101110 01111100 10001111"),
	}

	b := [][]byte{
		CreateBytesFromString("01101010 01011100 10001101"),
		CreateBytesFromString("011l1000 11010100 10l01101"),
		CreateBytesFromString("01001011 11111110 10101110"),
		CreateBytesFromString("01101110 01110100 10011101"),
		CreateBytesFromString("01101110 01110110 00001001"),
		CreateBytesFromString("11101110 01111100 10001111"),
	}

	printColumnln(t, "Original: ", a, 10)
	c := ToggleColumnSection(a, 10, 1, 4)
	assertMatrixMatches(t, b, c, 10, true)
}

func TestVerticalMultiColumnToggleSectionMatches(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("01101010 01011100 10001101"),
		CreateBytesFromString("011l1000 11110100 10l01101"),
		CreateBytesFromString("01001011 11011110 10101110"),
		CreateBytesFromString("01101110 01010100 10011101"),
		CreateBytesFromString("01101110 01010110 00001001"),
		CreateBytesFromString("11101110 01111100 10001111"),
	}

	b := [][]byte{
		CreateBytesFromString("01101010 01011100 10001101"),
		CreateBytesFromString("011l1001 00010100 10l01101"),
		CreateBytesFromString("01001010 00111110 10101110"),
		CreateBytesFromString("01101111 10110100 10011101"),
		CreateBytesFromString("01101111 10110110 00001001"),
		CreateBytesFromString("11101110 01111100 10001111"),
	}

	printColumnln(t, "Original: ", a, 10)
	c := ToggleColumnsSections(a, 7, 10, 1, 4)
	assertMatrixMatches(t, b, c, 9, true)
}

func TestFloatMask(t *testing.T) {
	a := createEmptyMatrix(1)
	b := [][]byte{
		CreateBytesFromString("1  0  1  1  1  1  1  0"),
		CreateBytesFromString("0  0  1  0  0  1  1  0"),
		CreateBytesFromString("1  0  1  0  0  1  1  0"),
		CreateBytesFromString("1  0  1  0  0  0  0  0"),
		CreateBytesFromString("0  0  1  1  1  1  1  1"),
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
	}
	c := ToggleUsingFloat64(a, 1701150464)

	assertMatrixMatchesPrintFull(t, b, c, true)
}

func TestUintMask(t *testing.T) {
	a := createEmptyMatrix(1)
	b := [][]byte{
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
		CreateBytesFromString("1  0  0  1  1  0  1  0"),
		CreateBytesFromString("1  0  0  1  1  0  1  0"),
		CreateBytesFromString("1  0  0  0  0  0  0  0"),
		CreateBytesFromString("1  1  1  1  1  1  1  1"),
	}
	c := ToggleUsingUint64(a, 1701150464)

	assertMatrixMatchesPrintFull(t, b, c, true)
}

func TestUintMaskReverse(t *testing.T) {
	a := createFullMatrix(1)
	b := [][]byte{
		CreateBytesFromString("0  0  0  0  0  0  0  0"),
		CreateBytesFromString("0  0  0  0  0  0  0  0"),
		CreateBytesFromString("0  0  0  0  0  0  0  0"),
		CreateBytesFromString("0  0  0  0  0  0  0  0"),
		CreateBytesFromString("0  1  1  0  0  1  0  1"),
		CreateBytesFromString("0  1  1  0  0  1  0  1"),
		CreateBytesFromString("0  1  1  1  1  1  1  1"),
		CreateBytesFromString("0  0  0  0  0  0  0  0"),
	}
	c := ToggleUsingUint64(a, 1701150464)

	assertMatrixMatchesPrintFull(t, b, c, true)
}

func assertByteSliceMatches(t *testing.T, a, b []byte) {
	for i, aBt := range a {
		if aBt != b[i] {
			t.Logf("%v - %v", aBt, b[i])
			t.Errorf("Byte slices should match:\n%v\n%v",
				SprintBytes(a), SprintBytes(b))
			return
		}
	}
}

func createEmptyMatrix(byteWidth int) [][]byte {
	matrix := make([][]byte, byteWidth*8)
	for i := range matrix {
		matrix[i] = make([]byte, byteWidth)
		for x := range matrix[i] {
			matrix[i][x] = byte(0)
		}
	}
	return matrix
}

func createFullMatrix(byteWidth int) [][]byte {
	matrix := make([][]byte, byteWidth*8)
	for i := range matrix {
		matrix[i] = make([]byte, byteWidth)
		for x := range matrix[i] {
			matrix[i][x] = byte(255)
		}
	}
	return matrix
}
