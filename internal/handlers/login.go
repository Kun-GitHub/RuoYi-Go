package handler

import (
	"RuoYi-Go/internal/response"
	"RuoYi-Go/pkg/captcha"
	ryredis "RuoYi-Go/pkg/redis"
	"fmt"
	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	var l loginStruct
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := ctx.ReadJSON(&l); err != nil {
		ctx.JSON(response.Error(iris.StatusBadRequest, fmt.Sprintf("Invalid JSON, error:%s", err.Error())))
		return

		//ctx.StopWithError(iris.StatusBadRequest, err)
		//return
	}

	v, error := ryredis.Redis.Get(fmt.Sprintf("captcha:%d", l.Uuid))
	if error != nil || v == "" {
		ctx.JSON(response.Error(iris.StatusInternalServerError, "验证码错误或已失效"))
		return
	}

	if v != "" && captcha.VerifyCaptcha(l.Uuid, l.Code) {

		//query.Use().SysUser.Table()

		ctx.WriteString("Hello, Iris!")
	} else {
		ctx.JSON(response.Error(iris.StatusInternalServerError, "验证码错误"))
		return
	}
}

type loginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}
