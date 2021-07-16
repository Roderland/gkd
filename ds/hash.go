package ds

type Hash interface {
	HSet(key, field, value string)
	HGet(key, field string) string
	HLen(key string) int
	HGetAll(key string) ([]string, []string)
	HDel(key string, fields [][]byte) int
	ToBytes() []byte
}
