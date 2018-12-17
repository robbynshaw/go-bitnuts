package bits

// SwapHorizontal takes a matrix of bytes and swaps a section from left to right.
func SwapHorizontal(matrix [][]byte, xStart, yStart, singleSectionWidth, height int) [][]byte {
	if int(singleSectionWidth) > (len(matrix)*8)/2 {
		panic("Single section width must be less or equal to than half the bits in a row")
	}

	if singleSectionWidth <= 0 {
		panic("Single section width must be greater than 0")
	}

	var fromRow []byte
	var toRow []byte
	newMatrix := CopyMatrix(matrix)
	toXStart := xStart + singleSectionWidth
	bitLen := len(matrix[0]) * 8
	isOverlapped := toXStart+singleSectionWidth >= bitLen

	for rowCount := 0; rowCount < height; rowCount++ {
		// If we went over the height of the matrix, wrap around to the top
		if yStart == len(matrix) {
			yStart = 0
		}

		fromRow = matrix[yStart]
		toRow = newMatrix[yStart]
		if isOverlapped {
			// If we went over the width of the matrix, wrap around to the left
			if toXStart > bitLen {
				// First section overlaps
				tempToXStart := toXStart - bitLen
				firstSecLen := bitLen - xStart
				secSecLen := singleSectionWidth - firstSecLen
				// Copy section to the end
				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, xStart, tempToXStart, firstSecLen)
				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, tempToXStart, xStart, firstSecLen)

				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, 0, tempToXStart+firstSecLen, secSecLen)
				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, tempToXStart+firstSecLen, 0, secSecLen)
			} else {
				// Second section overlaps
				firstSecLen := bitLen - toXStart
				secSecLen := singleSectionWidth - firstSecLen
				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, xStart, toXStart, firstSecLen)
				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, toXStart, xStart, firstSecLen)

				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, xStart+firstSecLen, 0, secSecLen)
				newMatrix[yStart] = CopyPasteRow(fromRow, toRow, 0, xStart+firstSecLen, secSecLen)
			}

		} else {
			newMatrix[yStart] = CopyPasteRow(fromRow, toRow, xStart, toXStart, singleSectionWidth)
			newMatrix[yStart] = CopyPasteRow(fromRow, toRow, toXStart, xStart, singleSectionWidth)
		}
		yStart++
	}

	return newMatrix
}

// SwapVertical takes a matrix of bytes and swaps a section from bottom to top.
func SwapVertical(matrix [][]byte, xStart, yStart, width, singleSectionHeight int) [][]byte {
	if int(singleSectionHeight)*2 > len(matrix) {
		panic("Single section height must be less than half the rows in the matrix")
	}

	if singleSectionHeight <= 0 {
		panic("Single section height must be greater than 0")
	}

	newMatrix := CopyMatrix(matrix)
	bitLen := len(matrix[0]) * 8
	toYStart := yStart + singleSectionHeight

	// Establish sections to copy
	secOneXStart := xStart
	secOneWidth := width
	secTwoXStart := 0
	secTwoWidth := 0
	if secOneXStart+secOneWidth > bitLen {
		secOneWidth = bitLen - secOneXStart
		secTwoXStart = 0
		secTwoWidth = width - secOneWidth
	}

	for rowCount := 0; rowCount < singleSectionHeight; rowCount++ {
		// If we went over the height of the matrix, wrap around to the top
		if yStart >= len(matrix) {
			yStart = yStart - len(matrix)
		}
		if toYStart >= len(matrix) {
			toYStart = toYStart - len(matrix)
		}

		newMatrix[toYStart] = CopyPasteRow(matrix[yStart], newMatrix[toYStart], secOneXStart, secOneXStart, secOneWidth)
		newMatrix[yStart] = CopyPasteRow(matrix[toYStart], newMatrix[yStart], secOneXStart, secOneXStart, secOneWidth)
		if secTwoWidth > 0 {
			newMatrix[toYStart] = CopyPasteRow(matrix[yStart], newMatrix[toYStart], secTwoXStart, secTwoXStart, secTwoWidth)
			newMatrix[yStart] = CopyPasteRow(matrix[toYStart], newMatrix[yStart], secTwoXStart, secTwoXStart, secTwoWidth)
		}
		yStart++
		toYStart++
	}

	return newMatrix
}
