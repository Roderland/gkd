package utils

import (
	"encoding/binary"
	"io"
)

type Message struct {
	Len  uint32
	Data string
}

func NewMessage(len uint32, data string) *Message {
	return &Message{
		Len:  len,
		Data: data,
	}
}

func NewMessageFromReader(reader io.Reader) (*Message, error) {
	bts := make([]byte, 4)
	_, err := reader.Read(bts)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(bts)
	bts = make([]byte, length)
	_, err = reader.Read(bts)
	if err != nil {
		return nil, err
	}
	return NewMessage(length, string(bts)), nil
}

func (m *Message) ToBytes() []byte {
	bts := make([]byte, m.Len+4)
	binary.BigEndian.PutUint32(bts[:4], m.Len)
	copy(bts[4:], m.Data)
	return bts
}
