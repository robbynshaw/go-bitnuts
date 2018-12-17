package bits

// Copy creates a new byte with the bits from start
// to finish copied and the rest of the bits as zeros
func Copy(from byte, start, finish uint8) byte {
	var i uint8
	retVal := uint8(from)
	for i = 0; i < 8; i++ {
		if i < start || i > finish {
			retVal = Unset(retVal, i)
		}
	}
	return byte(retVal)
}

// PasteRight pastes a result of Copy (or any byte section in
// which the rest of the byte contains only zeros) to the right
// side of byte 'to', starting at index 'start'
func PasteRight(from, to byte, start uint8) byte {
	to = to >> (8 - start)
	to = to << (8 - start)
	return to | from
}

// NewSimpleCopyPasteMask creates a bit mask in which
// one selected bit at 'index' is a 1 and the rest of
// the byte is composed of zeros.
func NewSimpleCopyPasteMask(index uint8) byte {
	mask := byte(1) << (7 - index)
	mask = ^mask
	return mask
}

// NewCopyPasteMask creates a bit mask in which the selected
// range is composed of 1's and the rest is composed
// of zeros. Both indexes are included.
func NewCopyPasteMask(start, stop uint8) byte {
	var i uint8
	b := byte(0)
	// TODO Make more efficient
	for i = 0; i < 8; i++ {
		if i >= start && i <= stop {
			b = Set(b, i)
		}
	}
	return b
}

// CopyPaste takes bit 'from' and copies all bits
// where 1's are found in the mask onto byte 'to'.
func CopyPaste(from, to, mask byte) byte {
	return from ^ (from^to)&mask
}

// CopyPasteRow takes all the bits of a defined section in a byte slice and pastes them
// onto an equal length section of another byte slice
func CopyPasteRow(fromRow, toRow []byte, fromStartIdx, toStartIdx, width int) []byte {
	fromEndIdx := fromStartIdx + width

	for fromStartIdx < fromEndIdx {
		if IsSetInRow(fromRow, fromStartIdx) {
			SetInRow(toRow, toStartIdx)
		} else {
			UnsetInRow(toRow, toStartIdx)
		}
		fromStartIdx++
		toStartIdx++
	}
	return toRow
}

// TODO This way is possibly more efficient but much harder
// func CopyPasteRow(fromRow, toRow []byte, fromStartIdx, toStartIdx, width int) []byte {
// 	var cp byte
// 	var partialCp byte
// 	var fromInByteStart uint8
// 	var fromInByteEnd uint8
// 	var toInByteStart uint8
// 	var toInByteEnd uint8
// 	byteIdx := 0
// 	endOfByte := 0
// 	tempEndIdx := 0
// 	pasteEndIdx := 0
// 	tempPasteEndIdx := 0
// 	rowBitWidth := len(fromRow) * 8
// 	fromEndIdx := fromStartIdx + width

// 	if len(fromRow) != len(toRow) {
// 		panic("Copy and Paste rows must be of equal length")
// 	}

// 	if fromStartIdx+width > rowBitWidth {
// 		panic("Start index + width must be less than or equal to the width of the byte slice (in bits)")
// 	}

// 	for bitIdx := 0; bitIdx < rowBitWidth; {
// 		byteIdx = bitIdx / 8
// 		endOfByte = ((byteIdx + 1) * 8) - 1

// 		if bitIdx >= fromStartIdx && bitIdx <= fromEndIdx {
// 			tempEndIdx = fromEndIdx
// 			if tempEndIdx > endOfByte {
// 				// copy the rest of the byte
// 				tempEndIdx = endOfByte
// 			}
// 			// TODO could be more efficient if we copied whole bytes, etc.
// 			fromInByteStart = uint8(bitIdx % 8)
// 			fromInByteEnd = uint8(tempEndIdx % 8)
// 			cp = Copy(fromRow[byteIdx], fromInByteStart, fromInByteEnd)

// 			// Paste
// 			pasteEndIdx = toStartIdx + (tempEndIdx - bitIdx)
// 			for toStartIdx <= pasteEndIdx {
// 				byteIdx = toStartIdx / 8
// 				endOfByte = ((byteIdx + 1) * 8) - 1
// 				tempPasteEndIdx = tempEndIdx
// 				if tempPasteEndIdx > endOfByte {
// 					// paste into the rest of the byte
// 					tempPasteEndIdx = endOfByte
// 				}

// 				// Shift bytes to the left to match pasting area
// 				partialCp = cp
// 				toInByteStart = uint8(toStartIdx % 8)
// 				if toInByteStart > fromInByteStart {
// 					partialCp <<= (toInByteStart - fromInByteStart)
// 				} else if toInByteStart < fromInByteStart {
// 					partialCp >>= (fromInByteStart - toInByteStart)
// 				}
// 				// clear out excess bits on the right
// 				for i := toInByteEnd + 1; i < 8; i++ {
// 					partialCp = Unset(partialCp, i)
// 				}

// 				toRow[byteIdx] = PasteRight(partialCp, toRow[byteIdx], toInByteStart)
// 				toStartIdx = tempPasteEndIdx
// 			}

// 			bitIdx = tempEndIdx + 1
// 		} else {
// 			bitIdx++
// 		}
// 	}
// 	return toRow
// }

// CopyMatrix creates a replica byte matrix
func CopyMatrix(in [][]byte) [][]byte {
	out := make([][]byte, len(in))

	for i := 0; i < len(in); i++ {
		out[i] = make([]byte, len(in[i]))
		for n := 0; n < len(in[i]); n++ {
			out[i][n] = in[i][n]
		}
	}
	return out
}
