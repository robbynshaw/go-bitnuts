package bits

import (
	"math"
)

// ShiftDown takes a matrix of bytes and shifts a column of bits down
// by 'count'. Bits pushed off the bottom are rotated back around to
// the top. Each slice in the matrix is expected to be the same length.
func ShiftDown(matrix [][]byte, index float64, count float64) [][]byte {
	yCursor := len(matrix) - int(count)
	return shiftVertical(matrix, index, count, yCursor)
}

// ShiftUp takes a matrix of bytes and shifts a column of bits up
// by 'count'. Bits pushed off the top are rotated back around to
// the bottom. Each slice in the matrix is expected to be the same length.
func ShiftUp(matrix [][]byte, index float64, count float64) [][]byte {
	yCursor := int(count)
	return shiftVertical(matrix, index, count, yCursor)
}

func shiftVertical(matrix [][]byte, index float64, count float64, yCursor int) [][]byte {
	validateVerticalShift(matrix, index, count)

	xByte, mask := getMaskForByteArrayIndex(index)

	froms := copyColumnBytes(matrix, yCursor, xByte)
	copyIntoMatrix(froms, matrix, xByte, mask)
	return matrix
}

func validateVerticalShift(matrix [][]byte, index float64, count float64) {
	if int(count) >= len(matrix) {
		panic("Must shift by a number less than the bits in the array")
	}

	if !matrixIsLongEnough(matrix, index) {
		panic("Each slice in the matrix must be as long as the index")
	}

	if count <= 0 {
		panic("Must shift by a number greater than 0")
	}
}

func matrixIsLongEnough(matrix [][]byte, index float64) bool {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) < (int(index)/8 + 1) {
			return false
		}
	}
	return true
}

func getMaskForByteArrayIndex(index float64) (xByte int, mask byte) {
	div := index / 8.0
	xByte = int(math.Trunc(div))

	offset := math.Remainder(index, 8)
	if offset < 0 {
		offset += 8
	}
	mask = NewSimpleCopyPasteMask(uint8(offset))
	return
}

func copyColumnBytes(matrix [][]byte, rowStart, xIndex int) []byte {
	froms := make([]byte, len(matrix))
	for i := 0; i < len(matrix); i++ {
		froms[i] = matrix[rowStart][xIndex]
		rowStart++
		if rowStart >= len(matrix) {
			rowStart = 0
		}
	}
	return froms
}

func copyIntoMatrix(from []byte, into [][]byte, xIndex int, mask byte) {
	for i := 0; i < len(into); i++ {
		into[i][xIndex] = CopyPaste(from[i], into[i][xIndex], mask)
	}
}
