package handler

import (
	"RuoYi-Go/internal/middlewares"
	"RuoYi-Go/internal/models"
	"RuoYi-Go/internal/responses"
	"RuoYi-Go/internal/services"
	"RuoYi-Go/pkg/logger"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

func GetRouters(ctx iris.Context) {
	loginUser := middlewares.GetLoginUser()
	if loginUser == nil || loginUser.UserID == 0 {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	var menus = make([]models.SysMenu, 0)
	var err error
	if loginUser.Admin {
		menus, err = services.GetAllMenus()
		if err != nil {
			logger.Log.Error("getRouters error,", zap.Error(err))
			ctx.JSON(responses.Error(iris.StatusInternalServerError, "获取菜单失败"))
			return
		}
	}

	// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
	ctx.JSON(responses.Success(menus))
}
