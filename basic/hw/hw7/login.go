package hw7

import (
	"context"
	"encoding/json"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

type LoginApi interface {
	LoginSMS(ctx *gin.Context)
}

type MyUserHandler struct {
	Svc         UserService
	codeSvc     CodeService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	Handler
	Cmd redis.Cmdable
	mdl []Middleware
}

func (u *MyUserHandler) LoginSMS(ctx *gin.Context) {
	root := u.loginSMS
	for i := len(u.mdl) - 1; i >= 0; i-- {
		m := u.mdl[i]
		root = m(root)
	}

	//	这里还可以结合使用middleware机制来处理
	sms := root(ctx)
	if !sms.IsSuccess() {
		// 这里具体按照错误码值再处理错误
		ctx.String(sms.GetCode(), sms.GetMsg())
	} else {
		ctx.String(sms.GetCode(), sms.GetT())
	}
}

func (u *MyUserHandler) loginSMS(ctx *gin.Context) *Res[string] {
	var req User
	if err := ctx.Bind(&req); err != nil {
		return &Res[string]{
			code:    404,
			msg:     err.Error(),
			success: false,
		}
	}
	login, _ := u.Svc.Login(context.Background(), req.Email, "123")

	marshal, _ := json.Marshal(login)

	return &Res[string]{
		code:    200,
		msg:     "操作成功",
		success: true,
		data:    string(marshal),
	}
}

type Res[T any] struct {
	code    int
	msg     string
	success bool
	data    T
}

func (r *Res[T]) GetCode() int {
	return r.code
}

func (r *Res[T]) GetMsg() string {
	return r.msg
}

func (r *Res[T]) IsSuccess() bool {
	return r.success
}

func (r *Res[T]) GetT() T {
	return r.data
}

type Result[T any] interface {
	GetCode() int
	GetMsg() string
	IsSuccess() bool
	GetT() T
}

type HandleFunc func(ctx *gin.Context) *Res[string]

type Middleware func(next HandleFunc) HandleFunc
