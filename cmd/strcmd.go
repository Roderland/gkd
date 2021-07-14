package cmd

import "gkd/db"

func Set(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	return db.Set(args[0], args[1])
}

func Get(db *db.Database, args []string) string {
	if len(args) < 1 {
		return "(error) wrong number of arguments"
	}
	return db.Get(args[0])
}

func SetNx(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	return db.SetNx(args[0], args[1])
}

func GetSet(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	return db.GetSet(args[0], args[1])
}

func StrLen(db *db.Database, args []string) string {
	if len(args) < 1 {
		return "(error) wrong number of arguments"
	}
	return db.StrLen(args[0])
}

func Append(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	return db.Append(args[0], args[1])
}

func IncrBy(db *db.Database, args []string) string {
	if len(args) < 2 {
		return "(error) wrong number of arguments"
	}
	return db.IncrBy(args[0], args[1])
}
