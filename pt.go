package pt

import "encoding/binary"

type coordKey [14]byte

var (
	coordPrefix = []byte("c_")
)

func KeyBigEndian() (key coordKey) {
	copy(key[:], coordPrefix)
	binary.BigEndian.PutUint32(key[2:6], 12345567)
	binary.BigEndian.PutUint64(key[6:14], 234253231)
	return
}

func KeyLittleEndian() (key coordKey) {
	copy(key[:], coordPrefix)
	binary.LittleEndian.PutUint32(key[2:6], 12345567)
	binary.LittleEndian.PutUint64(key[6:14], 234253231)
	return
}
