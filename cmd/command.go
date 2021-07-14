package cmd

import "gkd/db"

type CommandFunction func(*db.Database, []string) string

var CommandMap = make(map[string]CommandFunction)

func RegisterCmd(cmd string, cmdFunc CommandFunction) {
	CommandMap[cmd] = cmdFunc
}

func TestCmdFunc(db *db.Database, args []string) string {
	return "hello gkd"
}

func Save(db *db.Database, args []string) string {
	db.SaveData(".", "test.data")
	return "OK"
}

func init() {
	RegisterCmd("test", TestCmdFunc)
	RegisterCmd("save", Save)
	/* str command */
	RegisterCmd("set", Set)
	RegisterCmd("get", Get)
	RegisterCmd("setnx", SetNx)
	RegisterCmd("getset", GetSet)
	RegisterCmd("strlen", StrLen)
	RegisterCmd("append", Append)
	RegisterCmd("incrby", IncrBy)
	/* list command */
	RegisterCmd("lpush", LPush)
	RegisterCmd("rpush", RPush)
	RegisterCmd("lrange", LRange)
	RegisterCmd("llen", LLen)
	RegisterCmd("lpop", LPop)
	RegisterCmd("lindex", LIndex)
	RegisterCmd("rpop", RPop)
	RegisterCmd("lrem", LRem)
	RegisterCmd("lset", LSet)
}
