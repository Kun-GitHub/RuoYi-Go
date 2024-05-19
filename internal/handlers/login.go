package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/model"
	"RuoYi-Go/internal/response"
	rydb "RuoYi-Go/pkg/db"
	ryjwt "RuoYi-Go/pkg/jwt"
	ryredis "RuoYi-Go/pkg/redis"
	"fmt"
	"github.com/kataras/iris/v12"
	"strings"
)

func Login(ctx iris.Context) {
	var l loginStruct
	// Attempt to read and bind the JSON request body to the 'user' variable
	if err := ctx.ReadJSON(&l); err != nil {
		ctx.JSON(response.ErrorFormat(iris.StatusBadRequest, "Invalid JSON, error:%s", err.Error()))
		return
	}

	v, error := ryredis.Redis.Get(fmt.Sprintf("%s:%d", common.CAPTCHA, l.Uuid))
	if error != nil || v == "" {
		ctx.JSON(response.Error(iris.StatusInternalServerError, "验证码错误或已失效"))
		return
	}

	if v != "" && strings.EqualFold(v, l.Code) {
		sysUser := &model.SysUser{}

		if err := rydb.DB.FindColumns(model.TableNameSysUser, sysUser, "user_name = ? and status = '0'", l.Username); err != nil {
			ctx.JSON(response.Error(iris.StatusInternalServerError, "用户名或密码错误"))
			return
		}
		if sysUser.UserID == 0 || sysUser.Password != l.Password {
			ctx.JSON(response.Error(iris.StatusInternalServerError, "账号或密码错误"))
			return
		}

		token, error := ryjwt.Sign("user_id", fmt.Sprintf("%d", sysUser.UserID), 72)
		if error != nil {
			ctx.JSON(response.Error(iris.StatusInternalServerError, "生成token失败"))
		} else {
			user := loginSuccess{
				Code:    response.SUCCESS,
				Token:   token,
				Message: "操作成功",
			}
			// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
			ctx.JSON(user)
		}
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

type loginSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Token   string `json:"token"`
}
