package impl

import (
	"encoding/binary"
	"gkd/entry"
	"strconv"
)

type StrMap struct {
	obj map[string]string
}

func NewStrMap() *StrMap {
	return &StrMap{
		obj: make(map[string]string),
	}
}

func (sm *StrMap) Set(key, value string) {
	sm.obj[key] = value
}

func (sm *StrMap) SetNx(key, value string) int {
	if sm.Get(key) != "" {
		return 0
	}
	sm.Set(key, value)
	return 1
}

func (sm *StrMap) Get(key string) string {
	return sm.obj[key]
}

func (sm *StrMap) GetSet(key, value string) string {
	old := sm.Get(key)
	sm.Set(key, value)
	return old
}

func (sm *StrMap) StrLen(key string) int {
	return len(sm.Get(key))
}

func (sm *StrMap) Append(key, value string) int {
	value = sm.Get(key) + value
	sm.Set(key, value)
	return len(value)
}

func (sm *StrMap) IncrBy(key, increment string) (string, error) {
	in, err := strconv.Atoi(increment)
	if err != nil {
		return "increment is not a integer", err
	}
	va, err := strconv.Atoi(sm.Get(key))
	if err != nil {
		return "value is not a integer", err
	}
	sum := va + in
	if (in > 0 && sum < va) || (in < 0 && sum > va) {
		return "result is out of range", err
	}
	value := strconv.Itoa(sum)
	sm.Set(key, value)
	return value, nil
}

func (sm *StrMap) ToBytes() []byte {
	var offset int64 = 0
	var buf []byte
	for k, v := range sm.obj {
		tmp := make([]byte, 10+int64(len(k))+int64(len(v)))
		buf = append(buf, tmp...)
		binary.BigEndian.PutUint16(buf[offset:offset+2], entry.StrMark)
		offset += 2
		binary.BigEndian.PutUint32(buf[offset:offset+4], uint32(len(k)))
		offset += 4
		n := copy(buf[offset:], k)
		offset += int64(n)
		binary.BigEndian.PutUint32(buf[offset:offset+4], uint32(len(v)))
		offset += 4
		n = copy(buf[offset:], v)
		offset += int64(n)
	}
	return buf[:len(buf)]
}
