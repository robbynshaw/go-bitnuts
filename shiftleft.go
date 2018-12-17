package bits

import (
	"math"
)

// ShiftLeft takes an array of bytes lined up horizontally, bit to bit,
// and shifts the whole array left by 'count'. Bits pushed off the left
// edge are rotated to the right side.
func ShiftLeft(bytes []byte, count float64) []byte {
	if int(count) >= (len(bytes) * 8) {
		panic("Must shift by a number less than the bits in the array")
	}

	if count <= 0 {
		panic("Must shift by a number greater than 0")
	}

	div := count / 8.0
	start := math.Trunc(div)

	if div == start {
		return shiftLeftEven(bytes, int(count), int(start))
	}
	return shiftLeftUneven(bytes, int(count), int(start))
}

func shiftLeftEven(bytes []byte, count int, byteCursor int) []byte {
	newBytes := make([]byte, len(bytes))

	for i := 0; i < len(bytes); i++ {
		newBytes[i] = bytes[byteCursor]

		byteCursor++
		if byteCursor >= len(bytes) {
			byteCursor = 0
		}
	}
	return newBytes
}

func shiftLeftUneven(bytes []byte, count int, byteCursor int) []byte {
	var from byte
	newBytes := make([]byte, len(bytes))
	fltRmndr := math.Remainder(float64(count), float64(8))
	if fltRmndr < 0 {
		fltRmndr += 8
	}
	rmndr := uint8(fltRmndr)
	rmndrOp := 8 - rmndr
	isStart := true

	for i := 0; i < len(bytes); {
		if isStart {
			newBytes[i] = byte(0)
			from = Copy(bytes[byteCursor], rmndr, 8)
			from = from << rmndr
			newBytes[i] |= from

			byteCursor++
			if byteCursor >= len(bytes) {
				byteCursor = 0
			}
		} else {
			from = Copy(bytes[byteCursor], 0, rmndr-1)
			from = from >> rmndrOp
			newBytes[i] |= from
			i++
		}
		isStart = !isStart
	}
	return newBytes
}
