package impl

import (
	"encoding/binary"
	"gkd/entry"
	"sync"
)

type HashMap struct {
	obj   map[string]map[string]string
	mutex sync.RWMutex
}

func NewHashMap() *HashMap {
	return &HashMap{
		obj: make(map[string]map[string]string),
	}
}

func (hm *HashMap) HSet(key, field, value string) {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	if hm.obj[key] == nil {
		hm.obj[key] = make(map[string]string)
	}
	hm.obj[key][field] = value
}

func (hm *HashMap) HGet(key, field string) string {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()
	if hm.obj[key] == nil {
		return ""
	}
	return hm.obj[key][field]
}

func (hm *HashMap) HLen(key string) int {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()
	return len(hm.obj[key])
}

func (hm *HashMap) HGetAll(key string) ([]string, []string) {
	hm.mutex.RLock()
	defer hm.mutex.RUnlock()
	index := 0
	l := len(hm.obj[key])
	fields := make([]string, l)
	values := make([]string, l)
	for f, v := range hm.obj[key] {
		fields[index] = f
		values[index] = v
		index++
	}
	return fields, values
}

func (hm *HashMap) HDel(key string, fields [][]byte) int {
	hm.mutex.Lock()
	defer hm.mutex.Unlock()
	m := hm.obj[key]
	if m == nil {
		return 0
	}
	oldLen := len(m)
	for _, field := range fields {
		delete(m, string(field))
	}
	return oldLen - len(m)
}

func (hm *HashMap) ToBytes() []byte {
	var offset int64 = 0
	var buf []byte
	var tmp []byte
	for k, v := range hm.obj {
		tmp = make([]byte, 6+len(k)+4)
		buf = append(buf, tmp...)
		binary.BigEndian.PutUint16(buf[offset:offset+2], entry.HashMark)
		offset += 2
		binary.BigEndian.PutUint32(buf[offset:offset+4], uint32(len(k)))
		offset += 4
		n := copy(buf[offset:], k)
		offset += int64(n)
		binary.BigEndian.PutUint32(buf[offset:offset+4], uint32(len(v)))
		offset += 4
		for field, value := range v {
			tmp = make([]byte, 8+len(field)+len(value))
			buf = append(buf, tmp...)
			binary.BigEndian.PutUint32(buf[offset:offset+4], uint32(len(field)))
			offset += 4
			n = copy(buf[offset:], field)
			offset += int64(n)
			binary.BigEndian.PutUint32(buf[offset:offset+4], uint32(len(value)))
			offset += 4
			n = copy(buf[offset:], value)
			offset += int64(n)
		}
	}
	return buf
}
