package converter

import (
	. "encoding/binary"
)

func BigEndianInt642Bytes(i int64) []byte {
	var buf = make([]byte, 8)
	BigEndian.PutUint64(buf, uint64(i))
	return buf
}
func LittelEndianInt642Bytes(i int64) []byte {
	var buf = make([]byte, 8)
	LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BigEndianBytes2Int64(bytes []byte) uint64 {
	// return uint64(BigEndian.Uint32(bytes))<<16 | uint64(LittleEndian.Uint16(bytes[4:]))
	return uint64(BigEndian.Uint64(bytes))
}
func LittleEndianBytes2Int64(bytes []byte) uint64 {
	// return uint64(BigEndian.Uint32(bytes))<<16 | uint64(LittleEndian.Uint16(bytes[4:]))
	return uint64(LittleEndian.Uint64(bytes))
}
