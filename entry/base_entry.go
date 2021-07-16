package entry

type BaseEntry struct {
	Mark    uint16
	KeySize uint32
	Key     []byte
}

func NewBaseEntry(mark uint16, keySize uint32, key []byte) *BaseEntry {
	return &BaseEntry{Mark: mark, KeySize: keySize, Key: key}
}

const (
	SizeOfMark      = 2
	SizeOfKeySize   = 4
	SizeOfValueSize = 4
	SizeOfListLen   = 4
)

const (
	StrMark uint16 = iota
	ListMark
	HashMark
)
