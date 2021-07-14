package ds

type Str interface {
	Set(key, value string)
	SetNx(key, value string) int
	Get(key string) string
	GetSet(key, value string) string
	StrLen(key string) int
	Append(key, value string) int
	IncrBy(key, increment string) (string, error)
	ToBytes() []byte
}
