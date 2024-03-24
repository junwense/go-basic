package hw7

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestWeb(t *testing.T) {
	r := gin.Default()

	r.POST("/users/profile", func(ctx *gin.Context) {
		var req UserInfo
		if err := ctx.Bind(&req); err != nil {
			return
		}

		ctx.JSON(200, req)

	})
	u := &MyUserHandler{}
	r.POST("/test", u.LoginSMS)
	r.Run(":8000")
}

type UserInfo struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	BirthDay string `json:"birth_day"`
	Desc     string `json:"desc"`
}
