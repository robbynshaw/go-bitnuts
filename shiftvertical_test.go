package bits

import (
	"fmt"
	"math"
	"testing"
)

func TestShiftDownPanicFromSmallMatrix(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}

	assertPanic(t, "Shift Down should fail if the slices aren't long enough", func() {
		_ = ShiftDown(a, 25, 5)
	})
}

func TestShiftDownPanicFromZeroShift(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}

	assertPanic(t, "Zero shift should fail", func() {
		_ = ShiftDown(a, 1, 0)
	})
}

func TestShiftDownPanicFromExessShift(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}

	assertPanic(t, "Excess shift should fail", func() {
		_ = ShiftDown(a, 10, 7)
	})
}

func TestShiftDownMatches_x12_y2(t *testing.T) {
	x := float64(12)
	y := float64(2)
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}
	// 101001 top to bottom

	b := [][]byte{
		CreateBytesFromString("011010100101010010001101"),
		CreateBytesFromString("011l10001111110010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101111000001001"),
		CreateBytesFromString("111011100101010010001111"),
	}
	// 011010 top to bottom

	printColumnln(t, "Original: ", a, x)
	c := ShiftDown(a, x, y)
	assertMatrixMatches(t, b, c, x, true)
}
func TestShiftDownMatches_x5_y4(t *testing.T) {
	x := float64(5)
	y := float64(4)
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}
	// 000111 top to bottom

	b := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l11001111010010l01101"),
		CreateBytesFromString("010011111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011010100101011000001001"),
		CreateBytesFromString("111010100101110010001111"),
	}
	// 011100 top to bottom

	printColumnln(t, "Original: ", a, x)
	c := ShiftDown(a, x, y)
	assertMatrixMatches(t, b, c, x, true)
}

func TestShiftDownMatches_x20_y1(t *testing.T) {
	x := float64(20)
	y := float64(1)
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000000001"),
		CreateBytesFromString("111011100101110010001111"),
	}
	// 111101 top to bottom

	b := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010000111"),
	}
	// 111110 top to bottom

	printColumnln(t, "Original: ", a, x)
	c := ShiftDown(a, x, y)
	assertMatrixMatches(t, b, c, x, true)
}

func TestShiftDownMatches_x16_y3(t *testing.T) {
	x := float64(16)
	y := float64(3)
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000000001"),
		CreateBytesFromString("111011100101110010001111"),
	}
	// 111101 top to bottom

	b := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010000l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011010000001"),
		CreateBytesFromString("111011100101110010001111"),
	}
	// 101111 top to bottom

	printColumnln(t, "Original: ", a, x)
	c := ShiftDown(a, x, y)
	assertMatrixMatches(t, b, c, x, true)
}

func TestShiftUpPanicFromSmallMatrix(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}

	assertPanic(t, "Shift Up should fail if the slices aren't long enough", func() {
		_ = ShiftUp(a, 25, 5)
	})
}

func TestShiftUpPanicFromZeroShift(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}

	assertPanic(t, "Zero shift should fail", func() {
		_ = ShiftUp(a, 1, 0)
	})
}

func TestShiftUpPanicFromExessShift(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}

	assertPanic(t, "Excess shift should fail", func() {
		_ = ShiftUp(a, 10, 7)
	})
}

func TestShiftUpAMatches(t *testing.T) {
	a := [][]byte{
		CreateBytesFromString("011010100101110010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101010010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}
	// 101001 top to bottom

	b := [][]byte{
		CreateBytesFromString("011010100101010010001101"),
		CreateBytesFromString("011l10001111010010l01101"),
		CreateBytesFromString("010010111101111010101110"),
		CreateBytesFromString("011011100101110010011101"),
		CreateBytesFromString("011011100101011000001001"),
		CreateBytesFromString("111011100101110010001111"),
	}
	// 001101 top to bottom

	printColumnln(t, "Original: ", a, 12)
	c := ShiftUp(a, 12, 3)
	assertMatrixMatches(t, b, c, 12, true)
}

func assertMatrixMatches(t *testing.T, expected, actual [][]byte, colToPrint float64, shouldMatch bool) {
	if len(expected) != len(actual) && shouldMatch {
		t.Errorf("Matrices are different lengths: %v - %v", len(expected), len(actual))
	}

	matches := true
	for i := 0; i < len(expected); i++ {
		for n := 0; n < len(expected[i]); n++ {
			if expected[i][n] != actual[i][n] {
				matches = false
			}
		}
	}
	if matches && !shouldMatch {
		t.Error("Matrices should not match")
	} else if !matches && shouldMatch {
		printColumnln(t, "Expected: ", expected, colToPrint)
		printColumnln(t, "Actual  : ", actual, colToPrint)
		t.Error("Matrices should match")
	}
}

func assertMatrixMatchesPrintFull(t *testing.T, expected, actual [][]byte, shouldMatch bool) {
	if len(expected) != len(actual) && shouldMatch {
		t.Errorf("Matrices are different lengths: %v - %v", len(expected), len(actual))
	}

	matches := true
	for i := 0; i < len(expected); i++ {
		for n := 0; n < len(expected[i]); n++ {
			if expected[i][n] != actual[i][n] {
				matches = false
			}
		}
	}
	if matches && !shouldMatch {
		t.Error("Matrices should not match")
	} else if !matches && shouldMatch {
		printMatrix(t, expected)
		t.Log("\n")
		printMatrix(t, actual)
		t.Error("Matrices should match")
	}
}

func printColumnln(t *testing.T, msg string, matrix [][]byte, index float64) {
	rmndr := int(index)
	if index > 7 {
		rmndr = int(index) % 8
	}

	btFlt := math.Trunc(index / 8)
	xByte := int(btFlt)
	s := ""
	for i := 0; i < len(matrix); i++ {
		bt := fmt.Sprintf("%08b", matrix[i][xByte])
		s += string(bt[int(rmndr)])
	}
	t.Logf("%v%v\n", msg, s)
}

func printMatrix(t *testing.T, matrix [][]byte) {
	for _, row := range matrix {
		ln := ""
		for _, col := range row {
			bt := fmt.Sprintf("%08b", col)
			for _, char := range bt {
				s := string(char)
				ln += fmt.Sprintf(" %v ", s)
			}
		}
		t.Log(ln)
	}
}
