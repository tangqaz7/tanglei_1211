package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tanglei_1211/messageboard/configs"
	"tanglei_1211/messageboard/lv1/console"
)

func Register(userinfo *gin.Context) bool {
	username := userinfo.PostForm("username")
	password := userinfo.PostForm("password")
	fmt.Println(username,password)
	user := &configs.User{
		ID:       3,
		PassWord: username,
		Name:     password,
		RID:      "0",
		RName:    "0",
		RMessage: "0",
		Good:     0,
		IsGood:   "0",
	}
	fmt.Println(user)
	stmt,errA := console.User()
	fmt.Printf(
		"准备用户数据库:%v",
		configs.Error(errA),
	)

	errB := console.Exec(stmt,user)
	fmt.Printf(
		"插入数据:%v",
		configs.Error(errB),
	)
	res := configs.Error(errB)
	return res
}



