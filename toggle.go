package bits

// Toggle the bit at index in a byte
func Toggle(bt byte, index uint8) byte {
	bt ^= byte(1) << (7 - index)
	return bt
}

// ToggleRow toggles all of the bits in a byte array
func ToggleRow(data []byte) []byte {
	for i := 0; i < len(data); i++ {
		data[i] = ^data[i]
	}
	return data
}

// ToggleRowSection toggles all of the bits from start to end bit index
// in a byte array
func ToggleRowSection(data []byte, start, end int) []byte {
	for byteIndex := range data {
		startOfByte := byteIndex * 8
		if start <= (startOfByte + 8) {
			if start <= startOfByte && end >= ((byteIndex+1)*8) {
				// Can toggle full byte
				data[byteIndex] = ^data[byteIndex]
			} else {
				for i := uint8(0); i < 8; i++ {
					if startOfByte >= start && startOfByte <= end {
						data[byteIndex] = Toggle(data[byteIndex], i)
					}
					startOfByte++
				}
			}
		}
	}
	return data
}

// ToggleRowsSections toggles all the bits in the byte array rows selected
// by start and end for a specific section
func ToggleRowsSections(data [][]byte, start, end, xStart, xEnd int) [][]byte {
	if start < 0 || end > len(data) {
		panic("Start and end are not in the correct range")
	}

	for ; start <= end; start++ {
		data[start] = ToggleRowSection(data[start], xStart, xEnd)
	}

	return data
}

// ToggleRows toggles all the bits in the byte array rows selected
// by start and end
func ToggleRows(data [][]byte, start, end int) [][]byte {
	if start < 0 || end > len(data) {
		panic("Start and end are not in the correct range")
	}

	for ; start <= end; start++ {
		data[start] = ToggleRow(data[start])
	}

	return data
}

// ToggleColumn toggles all the bits in the selected column of a
// byte matrix
func ToggleColumn(matrix [][]byte, index float64) [][]byte {
	xByte, mask := getMaskForByteArrayIndex(index)

	for i := 0; i < len(matrix); i++ {
		matrix[i][xByte] ^= (^mask)
	}
	return matrix
}

// ToggleColumnSection toggles all the bits in the selected column of a
// byte matrix from start to end
func ToggleColumnSection(matrix [][]byte, col float64, start, end int) [][]byte {
	xByte, mask := getMaskForByteArrayIndex(col)
	if start < 0 || end > (len(matrix)-1) {
		panic("Start and end are not in the correct range")
	}

	for ; start <= end; start++ {
		matrix[start][xByte] ^= (^mask)
	}
	return matrix
}

// ToggleColumnsSections toggles all the bits in the byte array columns selected
// by start and end for a specific section
func ToggleColumnsSections(matrix [][]byte, start, end float64, yStart, yEnd int) [][]byte {
	if start < 0 || end > float64((len(matrix[0])*8))-1 || yStart < 0 || yEnd > (len(matrix)*8)-1 {
		panic("Starts and ends are not in the correct range")
	}

	for ; start <= end; start++ {
		matrix = ToggleColumnSection(matrix, start, yStart, yEnd)
	}

	return matrix
}

// ToggleColumns toggles all the bits in the columns of the byte
// matrix selected by start and end
func ToggleColumns(data [][]byte, start, end int) [][]byte {
	if start < 0 || end > (len(data[0])*8) {
		panic("Start and end are not in the correct range")
	}

	for ; start <= end; start++ {
		data = ToggleColumn(data, float64(start))
	}

	return data
}

// ToggleUsingFloat64 takes the bytes in a given float and uses them as
// masks to toggle values across a matrix
func ToggleUsingFloat64(matrix [][]byte, f float64) [][]byte {
	masks := Float64ToBytes(f)
	return toggleUsingMasks(matrix, masks)
}

// ToggleUsingUint64 takes the bytes in a given uint and uses them as
// masks to toggle values across a matrix
func ToggleUsingUint64(matrix [][]byte, u uint64) [][]byte {
	masks := Uint64ToBytes(u)
	return toggleUsingMasks(matrix, masks)
}

func toggleUsingMasks(matrix [][]byte, masks []byte) [][]byte {
	var mask byte
	i := 0
	for y, row := range matrix {
		for x := range row {
			if i >= len(masks) {
				i = 0
			}
			mask = masks[i]
			matrix[y][x] ^= (^mask)
			i++
		}
	}
	return matrix
}
