package db

import (
	"strconv"
)

func (db *Database) Set(key string, value string) string {
	// entry := utils.NewEntry([]byte(key), []byte(value), utils.PUT)

	db.strDict.Set(key, value)
	return "OK"
}

func (db *Database) Get(key string) string {
	value := db.strDict.Get(key)
	if value == "" {
		return "(nil)"
	}
	return value
}

func (db *Database) SetNx(key, value string) string {
	nx := db.strDict.SetNx(key, value)
	return "(integer) " + strconv.Itoa(nx)
}

func (db *Database) GetSet(key, value string) string {
	res := db.strDict.GetSet(key, value)
	if res == "" {
		return "(nil)"
	}
	return res
}

func (db *Database) StrLen(key string) string {
	length := db.strDict.StrLen(key)
	return "(integer) " + strconv.Itoa(length)
}

func (db *Database) Append(key, value string) string {
	length := db.strDict.Append(key, value)
	return "(integer) " + strconv.Itoa(length)
}

func (db *Database) IncrBy(key, increment string) string {
	res, err := db.strDict.IncrBy(key, increment)
	if err != nil {
		return "(error) " + res
	}
	return "(integer) " + res
}
