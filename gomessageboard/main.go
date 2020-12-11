package main

import (
	"gomessageboard/cmd"
	"gomessageboard/dao"
)

func main() {
	dao.MysqlInit()
	cmd.Entrance()
}
