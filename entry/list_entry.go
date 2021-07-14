package entry

type ListEntry struct {
	Base       *BaseEntry
	ListLen    uint32
	ValuesSize []uint32
	Values     [][]byte
}

func NewListEntry(base *BaseEntry, listLen uint32, valuesSize []uint32, values [][]byte) *ListEntry {
	return &ListEntry{Base: base, ListLen: listLen, ValuesSize: valuesSize, Values: values}
}
