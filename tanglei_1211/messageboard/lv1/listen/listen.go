package listen

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tanglei_1211/messageboard/lv1/handle"
)


func Register(userinfo *gin.Context) {
	res := handle.Register(userinfo)
	if res {
		userinfo.JSON(http.StatusOK, gin.H{
			"code":    2020,
			"message": "success",
		})
	} else {
		userinfo.JSON(http.StatusOK, gin.H{
			"code":    1010,
			"message": "failure",
		})
	}
}
