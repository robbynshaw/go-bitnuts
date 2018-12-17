package bits

// SpiralTopLeft toggles bits of a byte matrix in a spiral starting at the top left
// and going clockwise
func SpiralTopLeft(matrix [][]byte) [][]byte {
	// TODO Can't figure the dynamically sized version out yet

	// endCol := (len(matrix[0]) * 8) - 1
	// endRow := len(matrix)
	// row := -2
	// col := 0

	// for row+2 < endRow && col+2 < endCol {
	// 	// Right
	// 	row = row + 2
	// 	matrix[row] = ToggleRowSection(matrix[row], col, endCol-1)

	// 	// Down
	// 	endRow = endRow - 2
	// 	matrix = ToggleColumnSection(matrix, float64(endCol), row, endRow-1)

	// 	col = col + 2
	// 	// Left
	// 	matrix[endRow] = ToggleRowSection(matrix[endRow], col+1, endCol)

	// 	// Up
	// 	endCol = endCol + 2
	// 	matrix = ToggleColumnSection(matrix, float64(col), row+1, endRow)
	// }

	cols := (len(matrix[0]) * 8) - 1
	rows := len(matrix) - 1

	matrix[0] = ToggleRowSection(matrix[0], 0, cols)
	matrix = ToggleColumnSection(matrix, float64(cols), 1, cols-1)
	matrix[rows-1] = ToggleRowSection(matrix[rows-1], 1, cols-1)
	matrix = ToggleColumnSection(matrix, 1, 2, rows-2)
	matrix[2] = ToggleRowSection(matrix[2], 2, cols-2)
	matrix = ToggleColumnSection(matrix, float64(cols)-2, 3, rows-3)
	matrix[rows-3] = ToggleRowSection(matrix[rows-3], 3, cols-3)

	return matrix
}

// SpiralTopRight toggles bits of a byte matrix in a spiral starting at the top right
// and going clockwise
func SpiralTopRight(matrix [][]byte) [][]byte {
	// TODO Figure out dynamically sized

	cols := (len(matrix[0]) * 8) - 1
	rows := len(matrix) - 1

	matrix = ToggleColumnSection(matrix, float64(cols), 0, rows)
	matrix[rows] = ToggleRowSection(matrix[rows], 1, cols-1)
	matrix = ToggleColumnSection(matrix, 1, 2, rows-1)
	matrix[1] = ToggleRowSection(matrix[1], 1, cols-3)
	matrix = ToggleColumnSection(matrix, float64(cols-2), 1, rows-2)
	matrix[rows-2] = ToggleRowSection(matrix[rows-2], 3, cols-3)
	matrix = ToggleColumnSection(matrix, 3, 3, rows-3)

	return matrix
}

// SpiralBottomRight toggles bits of a byte matrix in a spiral starting at the bottom right
// and going clockwise
func SpiralBottomRight(matrix [][]byte) [][]byte {
	// TODO Figure out dynamically sized

	cols := (len(matrix[0]) * 8) - 1
	rows := len(matrix) - 1

	matrix[rows] = ToggleRowSection(matrix[rows], 0, cols)
	matrix = ToggleColumnSection(matrix, 0, 1, rows-1)
	matrix[1] = ToggleRowSection(matrix[1], 1, cols-1)
	matrix = ToggleColumnSection(matrix, float64(cols-1), 2, rows-2)
	matrix[rows-2] = ToggleRowSection(matrix[rows-2], 2, cols-2)
	matrix = ToggleColumnSection(matrix, 2, 3, rows-3)
	matrix[3] = ToggleRowSection(matrix[3], 3, cols-3)

	return matrix
}

// SpiralBottomLeft toggles bits of a byte matrix in a spiral starting at the bottom left
// and going clockwise
func SpiralBottomLeft(matrix [][]byte) [][]byte {
	// TODO Figure out dynamically sized

	cols := (len(matrix[0]) * 8) - 1
	rows := len(matrix) - 1

	matrix = ToggleColumnSection(matrix, 0, 0, rows)
	matrix[0] = ToggleRowSection(matrix[0], 1, cols-1)
	matrix = ToggleColumnSection(matrix, float64(cols-1), 1, rows-1)
	matrix[cols-1] = ToggleRowSection(matrix[cols-1], 2, cols-2)
	matrix = ToggleColumnSection(matrix, 2, 2, rows-2)
	matrix[2] = ToggleRowSection(matrix[2], 3, cols-3)
	matrix = ToggleColumnSection(matrix, float64(cols-3), 3, rows-3)

	return matrix
}
