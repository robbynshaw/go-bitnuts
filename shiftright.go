package bits

import (
	"math"
)

// ShiftRight takes an array of bytes lined up horizontally, bit to bit,
// and shifts the whole array right by 'count'. Bits pushed off the right
// edge are rotated to the left side.
func ShiftRight(bytes []byte, count float64) []byte {
	if int(count) >= (len(bytes) * 8) {
		panic("Must shift by a number less than the bits in the array")
	}

	if count <= 0 {
		panic("Must shift by a number greater than 0")
	}

	div := count / 8.0
	start := math.Trunc(div)
	isEven := div == start
	actStart := len(bytes) - 1 - int(start)

	if isEven {
		return shiftRightEven(bytes, int(count), int(actStart))
	}
	return shiftRightUneven(bytes, int(count), int(actStart))
}

func shiftRightEven(bytes []byte, count int, byteCursor int) []byte {
	newBytes := make([]byte, len(bytes))

	for i := len(bytes) - 1; i >= 0; i-- {
		newBytes[i] = bytes[byteCursor]

		byteCursor--
		if byteCursor < 0 {
			byteCursor = len(bytes) - 1
		}
	}
	return newBytes
}

func shiftRightUneven(bytes []byte, count int, byteCursor int) []byte {
	var from byte
	newBytes := make([]byte, len(bytes))
	fltRmndr := math.Remainder(float64(count), float64(8))
	if fltRmndr < 0 {
		fltRmndr += 8
	}
	rmndr := uint8(fltRmndr)
	rmndrOp := 8 - rmndr
	isStart := true

	for i := len(bytes) - 1; i >= 0; {
		if isStart {
			newBytes[i] = byte(0)
			from = Copy(bytes[byteCursor], 0, rmndrOp-1)
			from = from >> rmndr
			newBytes[i] |= from

			byteCursor--
			if byteCursor < 0 {
				byteCursor = len(bytes) - 1
			}
		} else {
			from = Copy(bytes[byteCursor], rmndrOp, 8)
			from = from << rmndrOp
			newBytes[i] |= from
			i--
		}
		isStart = !isStart
	}
	return newBytes
}
