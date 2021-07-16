package fr

import (
	"encoding/binary"
)

func (fr *FileReader) ReadHashLen() (hashLen uint32, err error) {
	buf := make([]byte, 4)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += 4
	hashLen = binary.BigEndian.Uint32(buf)
	return
}

func (fr *FileReader) ReadHashFieldSize() (fieldSize uint32, err error) {
	buf := make([]byte, 4)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += 4
	fieldSize = binary.BigEndian.Uint32(buf)
	return
}

func (fr *FileReader) ReadHashField(fieldSize uint32) (field []byte, err error) {
	buf := make([]byte, fieldSize)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += int64(fieldSize)
	field = buf
	return
}

func (fr *FileReader) ReadHashValueSize() (valueSize uint32, err error) {
	buf := make([]byte, 4)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += 4
	valueSize = binary.BigEndian.Uint32(buf)
	return
}

func (fr *FileReader) ReadHashValue(valueSize uint32) (value []byte, err error) {
	buf := make([]byte, valueSize)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += int64(valueSize)
	value = buf
	return
}
