package impl

import (
	"encoding/binary"
	"gkd/entry"
	"strconv"
	"sync"
)

type StrMap struct {
	obj   map[string]string
	mutex sync.RWMutex
}

func NewStrMap() *StrMap {
	return &StrMap{
		obj: make(map[string]string),
	}
}

func (sm *StrMap) Set(key, value string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.obj[key] = value
}

func (sm *StrMap) SetNx(key, value string) int {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	if sm.obj[key] != "" {
		return 0
	}
	sm.obj[key] = value
	return 1
}

func (sm *StrMap) Get(key string) string {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	return sm.obj[key]
}

func (sm *StrMap) GetSet(key, value string) string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	old := sm.obj[key]
	sm.obj[key] = value
	return old
}

func (sm *StrMap) StrLen(key string) int {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	return len(sm.obj[key])
}

func (sm *StrMap) Append(key, value string) int {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	value = sm.obj[key] + value
	sm.obj[key] = value
	return len(value)
}

func (sm *StrMap) IncrBy(key, increment string) (string, error) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	in, err := strconv.Atoi(increment)
	if err != nil {
		return "increment is not a integer", err
	}
	va, err := strconv.Atoi(sm.obj[key])
	if err != nil {
		return "value is not a integer", err
	}
	sum := va + in
	if (in > 0 && sum < va) || (in < 0 && sum > va) {
		return "result is out of range", err
	}
	value := strconv.Itoa(sum)
	sm.obj[key] = value
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
