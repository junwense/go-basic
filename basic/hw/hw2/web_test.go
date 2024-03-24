package hw2

import (
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestWeb(t *testing.T) {
	r := gin.Default()

	var birthDayRegexPattern = "^\\d{4}-\\d{2}-\\d{2}$"
	birthDay := regexp.MustCompile(birthDayRegexPattern, regexp.None)

	var userMap = make(map[int]*UserInfo)

	userMap[1] = &UserInfo{
		Id: 1,
	}

	r.POST("/users/edit", func(ctx *gin.Context) {

		var req UserInfo
		if err := ctx.Bind(&req); err != nil {
			return
		}
		info, ok := userMap[req.Id]

		if !ok {
			ctx.String(http.StatusOK, "当前用户不存在")
			return
		}

		if info == nil {
			ctx.String(http.StatusOK, "当前用户不存在")
			return
		}

		if req.BirthDay != "" {
			ok, err := birthDay.MatchString(req.BirthDay)

			if err != nil {
				ctx.String(http.StatusOK, "系统错误")
				return
			}
			if !ok {
				ctx.String(http.StatusOK, "你的生日格式不对")
				return
			}
		}

		err := validNikeName(req.NickName)

		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}
		err = validDesc(req.Desc)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		info.NickName = req.NickName
		info.BirthDay = req.BirthDay
		info.Desc = req.Desc

		ctx.JSON(200, "修改成功")
	})

	r.POST("/users/profile", func(ctx *gin.Context) {
		var req UserInfo
		if err := ctx.Bind(&req); err != nil {
			return
		}
		info, ok := userMap[req.Id]

		if !ok {
			ctx.String(http.StatusOK, "当前用户不存在")
			return
		}

		if info == nil {
			ctx.String(http.StatusOK, "当前用户不存在")
			return
		}

		ctx.JSON(200, info)

	})
	r.Run(":8000")
}

func validNikeName(name string) error {
	if len(name) > 10 {
		return errors.New("用户名字过长")
	}
	return nil
}

func validDesc(name string) error {
	if len(name) > 200 {
		return errors.New("备注过长")
	}
	return nil
}

// 昵称：字符串，你需要考虑允许的长度。
// • 生日：前端输入为 1992-01-01 这种字符串。
// • 个人简介：一段文本，你需要考虑允许的长度。
// • 尝试校验这些输入，并且返回准确的信息。
type UserInfo struct {
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	BirthDay string `json:"birth_day"`
	Desc     string `json:"desc"`
}
