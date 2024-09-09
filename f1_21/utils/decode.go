package utils

import (
	"encoding/binary"
	"math"
	"math/bits"
)

func D_uint8(b []byte) uint8 {
	return uint8(b[:1][0])
}

func D_int8(b []byte) int8 {
	return int8(b[:1][0])
}

func D_int16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(b))
}

func D_int32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(b))
}

func D_int64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}

func D_uint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func D_uint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func D_uint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func D_float(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func D_bool(b []byte) bool {
	return bits.OnesCount8(D_uint8(b)) != 0
}

func D_string(b []byte) string {
	return string(b)
}
