package ds

type List interface {
	LPush(key string, value ...[]byte) int
	RPush(key string, value [][]byte) int
	LRange(key string, start, end int) [][]byte
	LLen(key string) int
	LIndex(key string, index int) string
	LPop(key string) string
	RPop(key string) string
	LRem(key string, count int, value string) int
	LSet(key string, index int, value string) bool
	ToBytes() []byte
}
