package db

import "strconv"

func (db *Database) HSet(key, field, value string) string {
	db.hashDict.HSet(key, field, value)
	return "(integer) 1"
}

func (db *Database) HGet(key, field string) string {
	res := db.hashDict.HGet(key, field)
	if res == "" {
		return "(nil)"
	}
	return res
}

func (db *Database) HLen(key string) string {
	return "(integer) " + strconv.Itoa(db.hashDict.HLen(key))
}

func (db *Database) HGetAll(key string) string {
	fields, values := db.hashDict.HGetAll(key)
	l := len(fields)
	var res string
	for i := 0; i < l; i++ {
		if i != l-1 {
			res += fields[i] + " " + values[i] + "\n"
		} else {
			res += fields[i] + " " + values[i]
		}
	}
	return res
}

func (db *Database) HDel(key string, fields [][]byte) string {
	return "(integer) " + strconv.Itoa(db.hashDict.HDel(key, fields))
}
