package handler

import (
	"RuoYi-Go/internal/middlewares"
	"RuoYi-Go/internal/responses"
	"github.com/kataras/iris/v12"
)

func UserList(ctx iris.Context) {
	loginUser := middlewares.GetLoginUser()
	if loginUser == nil || loginUser.UserID == 0 {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

}
