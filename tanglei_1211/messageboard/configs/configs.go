package configs

import (
	"fmt"
)

//数据库连接信息
const (
	USERNAME = "mysql8"
	PASSWORD = "TangLei123"
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = 3306
	DATABASE = "MessageBoard"
)


var Connection = fmt.Sprintf(
	"%s:%s@%s(%s:%d)/%s?charset=utf8",
	USERNAME,
	PASSWORD,
	NETWORK,
	SERVER,
	PORT,
	DATABASE,
)
//信息
type Topic struct {
	ID int
	Name string
}
//用户信息
type User struct {
	ID int
	PassWord string
	Name string
	RID string
	RName string
	RMessage string
	Good int
	IsGood string
}

//login

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
//error
func Error(err error) bool {
	if err != nil {
		return false
	}else{
		return true
	}
}
