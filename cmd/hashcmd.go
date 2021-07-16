package cmd

import "gkd/db"

func HSet(db *db.Database, args []string) string {
	if len(args) < 3 {
		return "(error) wrong number of arguments"
	}
	return db.HSet(args[0], args[1], args[2])
}

func HGet(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	return db.HGet(args[0], args[1])
}

func HLen(db *db.Database, args []string) string {
	if len(args) < 1 {
		return "(error) wrong number of arguments"
	}
	return db.HLen(args[0])
}

func HGetAll(db *db.Database, args []string) string {
	if len(args) < 1 {
		return "(error) wrong number of arguments"
	}
	return db.HGetAll(args[0])
}

func HDel(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	var fields [][]byte
	for i := 1; i < len(args); i++ {
		fields = append(fields, []byte(args[i]))
	}
	return db.HDel(args[0], fields)
}
