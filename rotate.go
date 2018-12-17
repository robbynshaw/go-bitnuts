package bits

// TODO not currently working

// Rotate a byte matrix clockwise one turn
// WARNING: Not fast like the other methods in this package
// func Rotate(matrix [][]byte) [][]byte {
// 	bitLen := len(matrix[0]) * 8
// 	byteLen := len(matrix)
// 	r := make([][]byte, bitLen)
// 	buffer := byte(0)
// 	byteCursor := 0
// 	rowCursor := 0
// 	bitIx := uint8(0)
// 	for i := 0; i < byteLen; i++ { //byte in row
// 		for j := uint8(0); j < 8; j++ { // bit in byte
// 			r[rowCursor] = make([]byte, byteLen)
// 			for n := 0; n < bitLen; n++ { // row in matrix
// 				mask := NewSimpleCopyPasteMask(j)
// 				blank := byte(0)
// 				blank = CopyPaste(matrix[n][i], blank, mask)
// 				blank = blank << j
// 				blank = blank >> bitIx
// 				mask = NewSimpleCopyPasteMask(bitIx)
// 				buffer = CopyPaste(blank, buffer, mask)

// 				bitIx++
// 				if bitIx > 7 {
// 					bitIx = 0
// 					r[rowCursor][byteCursor] = buffer
// 					buffer = byte(0)
// 					byteCursor++
// 				}
// 				// Create byte from i in rows 0-7, etc.
// 			}
// 			rowCursor++
// 			byteCursor = 0
// 		}
// 	}
// 	return r
// }
