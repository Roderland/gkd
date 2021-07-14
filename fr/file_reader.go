package fr

import (
	"encoding/binary"
	"gkd/entry"
	"os"
)

type FileReader struct {
	File   *os.File
	Offset int64
}

func GetFile(path string, name string) (*os.File, error) {
	fileName := path + string(os.PathSeparator) + name
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewFileReader(path string, name string) (*FileReader, error) {
	file, err := GetFile(path, name)
	if err != nil {
		return nil, err
	}
	return &FileReader{
		File:   file,
		Offset: 0,
	}, nil
}

func (fr *FileReader) ReadMark() (mark uint16, err error) {
	buf := make([]byte, entry.SizeOfMark)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += entry.SizeOfMark
	mark = binary.BigEndian.Uint16(buf)
	return
}

func (fr *FileReader) ReadKeySize() (keySize uint32, err error) {
	buf := make([]byte, entry.SizeOfKeySize)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += entry.SizeOfKeySize
	keySize = binary.BigEndian.Uint32(buf)
	return
}

func (fr *FileReader) ReadKey(keySize uint32) (key []byte, err error) {
	buf := make([]byte, keySize)
	if _, err = fr.File.ReadAt(buf, fr.Offset); err != nil {
		return
	}
	fr.Offset += int64(keySize)
	key = buf
	return
}
