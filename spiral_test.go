package bits

import "testing"

func TestSpiralTopLeftCreatesPattern(t *testing.T) {
	blank := [][]byte{
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
	}

	fibbed := [][]byte{
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("0 0 0 0 0 0 0 1"),
		CreateBytesFromString("0 1 1 1 1 1 0 1"),
		CreateBytesFromString("0 1 0 0 0 1 0 1"),
		CreateBytesFromString("0 1 0 1 1 1 0 1"),
		CreateBytesFromString("0 1 0 0 0 0 0 1"),
		CreateBytesFromString("0 1 1 1 1 1 1 1"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
	}

	t.Log("Expected matrix:")
	printMatrix(t, fibbed)

	t.Log("Actual matrix:")
	newMatrix := SpiralTopLeft(blank)
	printMatrix(t, newMatrix)

	assertMatrixMatches(t, fibbed, newMatrix, 1, true)
}

func TestSpiralTopRightCreatesPattern(t *testing.T) {
	blank := [][]byte{
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
	}

	fibbed := [][]byte{
		CreateBytesFromString("0 0 0 0 0 0 0 1"),
		CreateBytesFromString("0 1 1 1 1 1 0 1"),
		CreateBytesFromString("0 1 0 0 0 1 0 1"),
		CreateBytesFromString("0 1 0 1 0 1 0 1"),
		CreateBytesFromString("0 1 0 1 0 1 0 1"),
		CreateBytesFromString("0 1 0 1 1 1 0 1"),
		CreateBytesFromString("0 1 0 0 0 0 0 1"),
		CreateBytesFromString("0 1 1 1 1 1 1 1"),
	}

	t.Log("Expected matrix:")
	printMatrix(t, fibbed)

	t.Log("Actual matrix:")
	newMatrix := SpiralTopRight(blank)
	printMatrix(t, newMatrix)

	assertMatrixMatches(t, fibbed, newMatrix, 1, true)
}

func TestSpiralBottomRightCreatesPattern(t *testing.T) {
	blank := [][]byte{
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
	}

	fibbed := [][]byte{
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("1 1 1 1 1 1 1 0"),
		CreateBytesFromString("1 0 0 0 0 0 1 0"),
		CreateBytesFromString("1 0 1 1 1 0 1 0"),
		CreateBytesFromString("1 0 1 0 0 0 1 0"),
		CreateBytesFromString("1 0 1 1 1 1 1 0"),
		CreateBytesFromString("1 0 0 0 0 0 0 0"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
	}

	t.Log("Expected matrix:")
	printMatrix(t, fibbed)

	t.Log("Actual matrix:")
	newMatrix := SpiralBottomRight(blank)
	printMatrix(t, newMatrix)

	assertMatrixMatches(t, fibbed, newMatrix, 1, true)
}

func TestSpiralBottomLeftCreatesPattern(t *testing.T) {
	blank := [][]byte{
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
		CreateBytesFromString("0 0 0 0 0 0 0 0"),
	}

	fibbed := [][]byte{
		CreateBytesFromString("1 1 1 1 1 1 1 0"),
		CreateBytesFromString("1 0 0 0 0 0 1 0"),
		CreateBytesFromString("1 0 1 1 1 0 1 0"),
		CreateBytesFromString("1 0 1 0 1 0 1 0"),
		CreateBytesFromString("1 0 1 0 1 0 1 0"),
		CreateBytesFromString("1 0 1 0 0 0 1 0"),
		CreateBytesFromString("1 0 1 1 1 1 1 0"),
		CreateBytesFromString("1 0 0 0 0 0 0 0"),
	}

	t.Log("Expected matrix:")
	printMatrix(t, fibbed)

	t.Log("Actual matrix:")
	newMatrix := SpiralBottomLeft(blank)
	printMatrix(t, newMatrix)

	assertMatrixMatches(t, fibbed, newMatrix, 1, true)
}
