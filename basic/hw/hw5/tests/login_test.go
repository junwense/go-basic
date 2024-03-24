package tests

import (
	"bytes"
	"encoding/json"
	"gitee.com/SeanJPK/go-basic/basic/hw/hw5"
	"gitee.com/SeanJPK/go-basic/basic/hw/hw5/mocks"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

// TestUserHandler_LoginSMS 由于循环依赖的问题，必须要移动到tests包下执行。
func TestUserHandler_LoginSMS(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	user := hw5.User{
		Id:    1,
		Email: "123@qq.com",
	}
	// http	测试使用此包
	resp := httptest.NewRecorder()
	tests := []struct {
		name string

		mockUserService func() hw5.UserService
		mockRedis       func() redis.Cmdable

		responseWriter http.ResponseWriter
		ctxFunc        func() *gin.Context
	}{
		{
			name:           "test",
			responseWriter: resp,
			ctxFunc: func() *gin.Context {
				resp = httptest.NewRecorder()
				marshal, _ := json.Marshal(user)
				s := string(marshal)
				bufferString := bytes.NewBufferString(s)
				req, _ := http.NewRequest(http.MethodPost, "/path", bufferString)
				req.Header.Set("Content-Type", "application/json")
				// 创建一个空的 gin.Context 对象
				ctx, _ := gin.CreateTestContext(resp)
				ctx.Request = req
				return ctx
			},

			mockUserService: func() hw5.UserService {
				service := hw5mocks.NewMockUserService(ctrl)
				service.EXPECT().Login(gomock.Any(), "123@qq.com", "123").Return(user, nil)
				return service
			},

			//mockRedis: func() redis.Cmdable {
			//	cmdable := hw5mocks.NewMockCmdable(ctrl)
			//	res := redis.NewCmd(context.Background(), nil)
			//	res.SetVal("OK")
			//	cmdable.EXPECT().Eval(gomock.Any(), "", []string{"locked-key"}, gomock.Any()).
			//		Return(res)
			//	return cmdable
			//},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &hw5.UserHandler{
				Svc: tt.mockUserService(),
				//codeSvc:     tt.fields.codeSvc,
				//emailExp:    tt.fields.emailExp,
				//passwordExp: tt.fields.passwordExp,
				//Handler:     tt.fields.Handler,
				//Cmd: tt.mockRedis(),
			}
			//u.LoginSMS(tt.args.ctx)
			// 初始化gin的ctx太麻烦了，这里可以采用流读取传入数据和传出数据进行校验
			// 可以定义一个对象，序列化成json和响应进行比较
			ctxFunc := tt.ctxFunc
			u.LoginSMS(ctxFunc())
			marshal, _ := json.Marshal(user)
			all, _ := ioutil.ReadAll(resp.Result().Body)
			result := string(all)
			assert.Equal(t, result, string(marshal))

		})
	}
}
