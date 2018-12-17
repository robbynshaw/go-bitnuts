package bits

// import (
// 	"testing"
// )

// func TestRotateMatches(t *testing.T) {
// 	a :=
// 		[][]byte{
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("1111111 11111111 111111111"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0000000 00000000 000000000"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("1111111 11111111 111111111"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0000000 00000000 000000000"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("1111111 11111111 111111111"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 			CreateBytesFromString("0111010 01000010 110111100"),
// 		}

// 	b := [][]byte{
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("11111110 11111101 11111111"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 		CreateBytesFromString("00001000 00100000 00100000"),
// 	}

// 	c := Rotate(a)

// 	assertMatrixMatches(t, b, c, 1, true)
// }
