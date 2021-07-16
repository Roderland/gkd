package entry

type HashEntry struct {
	Base       *BaseEntry
	HashLen    uint32
	FieldsSize []uint32
	Fields     [][]byte
	ValuesSize []uint32
	Values     [][]byte
}

func NewHashEntry(base *BaseEntry, hashLen uint32, fieldsSize []uint32, fields [][]byte,
	valuesSize []uint32, values [][]byte) *HashEntry {
	return &HashEntry{
		Base: base, HashLen: hashLen, FieldsSize: fieldsSize,
		Fields: fields, ValuesSize: valuesSize, Values: values,
	}
}
