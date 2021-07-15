package fr

import (
	"encoding/binary"
	"gkd/entry"
)

func (fr *FileReader) ReadListLen() (listLen uint32, err error) {
	buf := make([]byte, entry.SizeOfListLen)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += entry.SizeOfListLen
	listLen = binary.BigEndian.Uint32(buf)
	return
}

func (fr *FileReader) ReadListValueSize() (valueSize uint32, err error) {
	buf := make([]byte, entry.SizeOfValueSize)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += entry.SizeOfValueSize
	valueSize = binary.BigEndian.Uint32(buf)
	return
}

func (fr *FileReader) ReadListValue(valueSize uint32) (value []byte, err error) {
	buf := make([]byte, valueSize)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += int64(valueSize)
	value = buf
	return
}
