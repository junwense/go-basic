package hw5

import (
	"context"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

type UserHandler struct {
	Svc         UserService
	codeSvc     CodeService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	Handler
	Cmd redis.Cmdable
}

func (u *UserHandler) LoginSMS(ctx *gin.Context) {
	var req User
	if err := ctx.Bind(&req); err != nil {
		return
	}
	login, _ := u.Svc.Login(context.Background(), req.Email, "123")

	ctx.JSON(200, login)

}
