package entry

type StrEntry struct {
	Base      *BaseEntry
	ValueSize uint32
	Value     []byte
}

func NewStrEntry(base *BaseEntry, valueSize uint32, value []byte) *StrEntry {
	return &StrEntry{Base: base, ValueSize: valueSize, Value: value}
}
