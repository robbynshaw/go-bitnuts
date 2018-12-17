package bits

import (
	"encoding/binary"
	"math"
)

// Float64ToBytes converts a float64 into it's constituent bytes in big endian format
func Float64ToBytes(num float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(num))
	return buf[:]
}

// Uint64ToBytes converts a uint64 into it's constituent bytes in big endian format
func Uint64ToBytes(num uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], num)
	return buf[:]
}
