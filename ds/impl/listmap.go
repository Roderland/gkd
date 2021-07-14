package impl

import (
	"container/list"
)

type ListMap struct {
	obj map[string]*list.List
}

func NewListMap() *ListMap {
	return &ListMap{
		obj: make(map[string]*list.List),
	}
}

func (lm *ListMap) LPush(key string, value ...[]byte) int {
	var l *list.List
	if l = lm.obj[key]; l == nil {
		l = list.New()
		lm.obj[key] = l
	}
	for _, n := range value {
		l.PushFront(string(n))
	}
	return l.Len()
}

func (lm *ListMap) RPush(key string, value [][]byte) int {
	var l *list.List
	if l = lm.obj[key]; l == nil {
		l = list.New()
		lm.obj[key] = l
	}
	for _, n := range value {
		l.PushBack(string(n))
	}
	return l.Len()
}

func (lm *ListMap) LRange(key string, start, end int) (values [][]byte) {
	l := lm.obj[key]
	if start < 0 {
		start += l.Len()
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end += l.Len()
	}
	for index, p := 0, l.Front(); p != nil && index <= end; index, p = index+1, p.Next() {
		if index >= start {
			s := p.Value.(string)
			values = append(values, []byte(s))
		}
	}
	return
}

func (lm *ListMap) LLen(key string) int {
	l := lm.obj[key]
	if l == nil {
		return 0
	}
	return l.Len()
}

func (lm *ListMap) LIndex(key string, index int) (value string) {
	l := lm.obj[key]
	if index < 0 {
		index += l.Len()
		if index < 0 {
			return value
		}
	}
	if index >= l.Len() {
		return value
	}
	var i int
	var p *list.Element
	for i, p = 0, l.Front(); i != index; i, p = i+1, p.Next() {
	}
	return p.Value.(string)
}

func (lm *ListMap) LPop(key string) string {
	l := lm.obj[key]
	return l.Remove(l.Front()).(string)
}

func (lm *ListMap) RPop(key string) string {
	l := lm.obj[key]
	return l.Remove(l.Back()).(string)
}

func (lm *ListMap) LRem(key string, count int, value string) int {
	l := lm.obj[key]
	res := 0
	if count >= 0 {
		for p := l.Front(); p != nil && count > 0; {
			next := p.Next()
			if p.Value.(string) == value {
				l.Remove(p)
				res++
				count--
			}
			p = next
		}
	} else {
		count = -count
		for p := l.Back(); p != nil && count > 0; {
			prev := p.Prev()
			if p.Value.(string) == value {
				l.Remove(p)
				res++
				count--
			}
			p = prev
		}
	}
	return res
}

func (lm *ListMap) LSet(key string, index int, value string) bool {
	l := lm.obj[key]
	if index < 0 {
		index += l.Len()
		if index < 0 {
			return false
		}
	}
	if index >= l.Len() {
		return false
	}
	i, p := 0, l.Front()
	for ; i < index; i, p = i+1, p.Next() {
	}
	p.Value = value
	return true
}
