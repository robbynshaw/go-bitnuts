package bits

import (
	"testing"
)

func TestCreateGraphMaskSinZeros(t *testing.T) {
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
	graph := [][]byte{
		CreateBytesFromString("1 1 1 0 0 0 0 1"),
		CreateBytesFromString("1 0 1 0 0 0 1 1"),
		CreateBytesFromString("1 0 1 0 0 0 1 0"),
		CreateBytesFromString("1 0 1 1 0 0 1 0"),
		CreateBytesFromString("0 0 0 1 0 0 1 0"),
		CreateBytesFromString("0 0 0 1 0 1 0 0"),
		CreateBytesFromString("0 0 0 1 0 1 0 0"),
		CreateBytesFromString("0 0 0 0 1 1 0 0"),
	}

	t.Log("Expected matrix:")
	printMatrix(t, graph)

	t.Log("Actual matrix:")
	newMatrix := SinWave(blank, 4, 1)
	printMatrix(t, newMatrix)

	assertMatrixMatches(t, graph, newMatrix, 1, true)
}

func TestCreateGraphMaskSinOnes(t *testing.T) {
	blank := [][]byte{
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
		CreateBytesFromString("1 1 1 1 1 1 1 1"),
	}
	graph := [][]byte{
		CreateBytesFromString("0 0 0 1 1 1 1 0"),
		CreateBytesFromString("0 1 0 1 1 1 0 0"),
		CreateBytesFromString("0 1 0 1 1 1 0 1"),
		CreateBytesFromString("0 1 0 0 1 1 0 1"),
		CreateBytesFromString("1 1 1 0 1 1 0 1"),
		CreateBytesFromString("1 1 1 0 1 0 1 1"),
		CreateBytesFromString("1 1 1 0 1 0 1 1"),
		CreateBytesFromString("1 1 1 1 0 0 1 1"),
	}

	t.Log("Expected matrix:")
	printMatrix(t, graph)

	t.Log("Actual matrix:")
	newMatrix := SinWave(blank, 4, 1)
	printMatrix(t, newMatrix)

	assertMatrixMatches(t, graph, newMatrix, 1, true)
}
