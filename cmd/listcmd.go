package cmd

import (
	"gkd/db"
	"strconv"
)

func LPush(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	var values [][]byte
	for i := 1; i < len(args); i++ {
		values = append(values, []byte(args[i]))
	}
	return db.LPush(args[0], values...)
}

func RPush(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	var values [][]byte
	for i := 1; i < len(args); i++ {
		values = append(values, []byte(args[i]))
	}
	return db.RPush(args[0], values)
}

func LRange(db *db.Database, args []string) string {
	if len(args) < 3 {
		return "(error) wrong number of arguments"
	}
	start, err := strconv.Atoi(args[1])
	if err != nil {
		return "(error) argument start is not a integer"
	}
	end, err := strconv.Atoi(args[2])
	if err != nil {
		return "(error) argument end is not a integer"
	}
	return db.LRange(args[0], start, end)
}

func LLen(db *db.Database, args []string) string {
	if len(args) < 1 {
		return "(error) wrong number of arguments"
	}
	return db.LLen(args[0])
}

func LIndex(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	index, err := strconv.Atoi(args[1])
	if err != nil {
		return "(error) argument index is not a integer"
	}
	return db.LIndex(args[0], index)
}

func LPop(db *db.Database, args []string) string {
	if len(args) < 1 {
		return "(error) wrong number of arguments"
	}
	return db.LPop(args[0])
}

func RPop(db *db.Database, args []string) string {
	if len(args) < 1 {
		return "(error) wrong number of arguments"
	}
	return db.RPop(args[0])
}

func LRem(db *db.Database, args []string) string {
	if len(args) < 3 {
		return "(error) wrong number of arguments"
	}
	count, err := strconv.Atoi(args[1])
	if err != nil {
		return "(error) argument count is not a integer"
	}
	return db.LRem(args[0], count, args[2])
}

func LSet(db *db.Database, args []string) string {
	if len(args) < 3 {
		return "(error) wrong number of arguments"
	}
	index, err := strconv.Atoi(args[1])
	if err != nil {
		return "(error) argument index is not a integer"
	}
	return db.LSet(args[0], index, args[2])
}
