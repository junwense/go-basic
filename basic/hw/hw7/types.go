package hw7

import (
	"context"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(ctx context.Context, email, password string) (User, error)
	SignUp(ctx context.Context, u User) error
	FindOrCreate(ctx context.Context, phone string) (User, error)
	FindOrCreateByWechat(ctx context.Context, wechatInfo WechatInfo) (User, error)
	Profile(ctx context.Context, id int64) (User, error)
}

type CodeService interface {
	Send(ctx context.Context,
		// 区别业务场景
		biz string, phone string) error
	Verify(ctx context.Context, biz string,
		phone string, inputCode string) (bool, error)
}

type Handler interface {
	SetLoginToken(ctx *gin.Context, uid int64) error
	SetJWTToken(ctx *gin.Context, uid int64, ssid string) error
	ClearToken(ctx *gin.Context) error
	CheckSession(ctx *gin.Context, ssid string) error
	ExtractToken(ctx *gin.Context) string
}
