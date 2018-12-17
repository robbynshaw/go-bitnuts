package bits

import "math"

type grapher func(x float64) float64

// Graph takes f(x) and toggles the matching bits in a byte matrix
func Graph(matrix [][]byte, fx grapher) [][]byte {
	mask := createGraphMask(matrix, fx)
	return applyMask(matrix, mask)
}

// SinWave toggles a variable width sine wave on a byte matrix
func SinWave(matrix [][]byte, yStart, widthCoef float64) [][]byte {
	return Graph(matrix, func(x float64) float64 {
		return (4 * math.Sin(widthCoef*x)) + yStart
	})
}

// SinWaveFunky toggles a variable width sine wave on a byte matrix
func SinWaveFunky(matrix [][]byte, yStart, cosCoef, sinCoef float64) [][]byte {
	return Graph(matrix, func(x float64) float64 {
		return (yStart * math.Cos((cosCoef*x)/yStart) * math.Sin((sinCoef*x)/cosCoef)) + yStart
	})
}

func applyMask(matrix [][]byte, mask [][]byte) [][]byte {
	for y, maskRow := range mask {
		for x, maskByte := range maskRow {
			matrix[y][x] ^= maskByte
		}
	}
	return matrix
}

// TODO cache all graph masks
func createGraphMask(matrix [][]byte, fx grapher) [][]byte {
	var x float64
	masks := createMatrix(len(matrix), len(matrix[0])*8)

	for x = 0; x < float64(len(matrix[0])*8); x = x + 0.1 {
		y := len(matrix) - int(fx(x)) - 1
		// y := int(fx(x))
		byteIndex := int(x / 8)
		bitIndex := uint8(int(x) % 8)
		if y >= 0 && y < len(matrix) {
			masks[y][byteIndex] = Set(masks[y][byteIndex], bitIndex)
		}
	}
	return masks
}

func createMatrix(x, y int) [][]byte {
	matrix := make([][]byte, y)
	for i := range matrix {
		matrix[i] = make([]byte, x/8)
	}
	return matrix
}
