package bits

// Unset makes the bit at i in a byte 0
func Unset(bt byte, i uint8) byte {
	return bt &^ (byte(1) << (7 - i))
}

// UnsetInRow makes the bit at i in a slice of bits 0
func UnsetInRow(row []byte, i int) []byte {
	curByte := i / 8
	curBit := i % 8
	row[curByte] = Unset(row[curByte], uint8(curBit))
	return row
}

// Set makes the bit at i in a byte 1
func Set(bt byte, i uint8) byte {
	return bt | (byte(1) << (7 - i))
}

// SetInRow makes the bit at i in a slice of bits 1
func SetInRow(row []byte, i int) []byte {
	curByte := i / 8
	curBit := i % 8
	row[curByte] = Set(row[curByte], uint8(curBit))
	return row
}

// IsSet checks if the bit at i is set
func IsSet(bt byte, i uint8) bool {
	return bt&(byte(1)<<(7-i)) != byte(0)
}

// IsSetInRow checks if the bit at i in a slice of bits is set
func IsSetInRow(row []byte, i int) bool {
	curByte := i / 8
	curBit := i % 8
	return IsSet(row[curByte], uint8(curBit))
}
