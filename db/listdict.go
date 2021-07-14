package db

import "strconv"

func (db *Database) LPush(key string, value ...[]byte) string {
	length := db.listDict.LPush(key, value...)
	return "(integer) " + strconv.Itoa(length)
}

func (db *Database) RPush(key string, value [][]byte) string {
	length := db.listDict.RPush(key, value)
	return "(integer) " + strconv.Itoa(length)
}

func (db *Database) LRange(key string, start, end int) string {
	values := db.listDict.LRange(key, start, end)
	var res string
	for i, value := range values {
		res += string(value)
		if i != len(values)-1 {
			res += "\n"
		}
	}
	if res == "" {
		res = "(empty list)"
	}
	return res
}

func (db *Database) LLen(key string) string {
	return "(integer) " + strconv.Itoa(db.listDict.LLen(key))
}

func (db *Database) LIndex(key string, index int) string {
	res := db.listDict.LIndex(key, index)
	if res == "" {
		return "(nil)"
	}
	return res
}

func (db *Database) LPop(key string) string {
	res := db.listDict.LPop(key)
	if res == "" {
		return "(nil)"
	}
	return res
}

func (db *Database) RPop(key string) string {
	res := db.listDict.RPop(key)
	if res == "" {
		return "(nil)"
	}
	return res
}

func (db *Database) LRem(key string, count int, value string) string {
	rem := db.listDict.LRem(key, count, value)
	return "(integer) " + strconv.Itoa(rem)
}

func (db *Database) LSet(key string, index int, value string) string {
	res := db.listDict.LSet(key, index, value)
	if !res {
		return "(error) index out of range"
	}
	return "OK"
}
