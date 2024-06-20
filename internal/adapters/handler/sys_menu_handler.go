// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"github.com/kataras/iris/v12"
	"net/url"
	"strings"
)

type SysMenuHandler struct {
	service input.SysMenuService
}

func NewSysMenuHandler(service input.SysMenuService) *SysMenuHandler {
	return &SysMenuHandler{service: service}
}

// GenerateCaptchaImage
func (h *SysMenuHandler) GetRouters(ctx iris.Context) {
	user := ctx.Values().Get(common.LOGINUSER)
	// 类型断言
	loginUser, ok := user.(*model.LoginUserStruct)
	if !ok {
		ctx.JSON(common.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	var menus = make([]*model.SysMenu, 0)
	menus, err := h.service.QueryMenusByUserId(loginUser.UserID)
	if err != nil {
		//h.logger.Debug("login failed", zap.Error(err))
		ctx.JSON(common.ErrorFormat(iris.StatusInternalServerError, "GetRouters, error：%s", err.Error()))
		return
	}
	ctx.JSON(common.Success(buildMenuTree(menus)))
}

type MetaStruct struct {
	Title   string `json:"title,omitempty"`
	Icon    string `json:"icon,omitempty"`
	NoCache bool   `json:"noCache"`
	//Link    string `json:"link,omitempty"`
}

type routerStruct struct {
	Hidden     bool            `json:"hidden"`
	AlwaysShow bool            `json:"alwaysShow"`
	Name       string          `json:"name,omitempty"`
	Path       string          `json:"path,omitempty"`
	Component  string          `json:"component,omitempty"`
	Redirect   string          `json:"redirect,omitempty"`
	Query      string          `json:"query,omitempty"`
	Meta       *MetaStruct     `json:"meta,omitempty"`
	Children   []*routerStruct `json:"children,omitempty"`
}

func buildMenuTree(menus []*model.SysMenu) []*routerStruct {
	menuMap := make(map[int64]*routerStruct)
	rootMenus := make([]*routerStruct, 0)

	for _, menu := range menus {
		if menu.ParentID == 0 {
			router := &routerStruct{
				Hidden:    menu.Visible == "1",
				Name:      getRouteName(menu),
				Path:      getRouterPath(menu),
				Component: getComponent(menu),
				Redirect: func() string {
					if menu.MenuType == common.TYPE_DIR {
						return "noRedirect"
					}
					return ""
				}(),
				AlwaysShow: func() bool {
					if menu.MenuType == common.TYPE_DIR {
						return true
					}
					return false
				}(),
				Query: menu.Query,
				Meta: &MetaStruct{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: menu.IsCache == 1,
				},
			}

			menuMap[menu.MenuID] = router
			rootMenus = append(rootMenus, router)
		} else {
			if parent, ok := menuMap[menu.ParentID]; ok {
				router := &routerStruct{
					Hidden:    menu.Visible == "1",
					Name:      getRouteName(menu),
					Path:      getRouterPath(menu),
					Component: getComponent(menu),
					Redirect: func() string {
						if menu.MenuType == common.TYPE_DIR {
							return "noRedirect"
						}
						return ""
					}(),
					AlwaysShow: func() bool {
						if menu.MenuType == common.TYPE_DIR {
							return true
						}
						return false
					}(),
					Query: menu.Query,
					Meta: &MetaStruct{
						Title:   menu.MenuName,
						Icon:    menu.Icon,
						NoCache: menu.IsCache == 1,
					},
				}

				if parent.Children == nil {
					parent.Children = make([]*routerStruct, 0)
				}
				parent.Children = append(parent.Children, router)
				menuMap[menu.MenuID] = router
			} else {

			}
		}
	}
	return rootMenus
}

// 注意：以下辅助函数需要根据实际情况实现
func getRouteName(menu *model.SysMenu) string {
	routerName := strings.Title(menu.Path)
	// Non-outer link and is a first-level directory (type is directory)
	if isMenuFrame(menu) {
		routerName = ""
	}
	return routerName
}

func getRouterPath(menu *model.SysMenu) string {
	routerPath := menu.Path

	// Inner link open external way
	if menu.ParentID != 0 && isInnerLink(menu) {
		routerPath = innerLinkReplaceEach(routerPath)
	}

	// Not an outer link and is a top-level directory (type is directory)
	if menu.ParentID == 0 && menu.MenuType == common.TYPE_DIR && menu.IsFrame == common.NO_FRAME {
		routerPath = "/" + menu.Path
	} else if isMenuFrame(menu) {
		routerPath = "/"
	}
	return routerPath
}

func getComponent(menu *model.SysMenu) string {
	component := common.LAYOUT
	if strings.TrimSpace(menu.Component) != "" && !isMenuFrame(menu) {
		component = menu.Component
	} else if strings.TrimSpace(menu.Component) == "" && menu.ParentID != 0 && isInnerLink(menu) {
		component = common.INNER_LINK
	} else if strings.TrimSpace(menu.Component) == "" && isParentView(menu) {
		component = common.PARENT_VIEW
	}
	return component
}

func isMenuFrame(menu *model.SysMenu) bool {
	return menu.ParentID == 0 && common.TYPE_MENU == menu.MenuType && menu.IsFrame == common.NO_FRAME
}

func isInnerLink(menu *model.SysMenu) bool {
	return menu.IsFrame == common.NO_FRAME && isHTTP(menu.Path)
}

// isParentView checks if the given menu is a parent view.
// This function needs to be implemented based on your specific logic.
func isParentView(menu *model.SysMenu) bool {
	// Implement your logic here.
	return menu.ParentID != 0 && common.TYPE_DIR == menu.MenuType
}

// isHTTP checks if the provided string is a valid HTTP or HTTPS URL.
func isHTTP(urlStr string) bool {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return false
	}
	return parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
}

func innerLinkReplaceEach(path string) string {
	// 实现逻辑
	// 创建一个替换器，用于替换path中的目标子串
	replacer := strings.NewReplacer(
		HTTP, "",
		HTTPS, "",
		WWW, "",
		DOT, SLASH,
		COLON, SLASH,
	)
	// 使用替换器替换path中的目标子串
	return replacer.Replace(path)
}

const (
	HTTP  = "http://"
	HTTPS = "https://"
	WWW   = "www."
	DOT   = "."
	COLON = ":"
	SLASH = "/"
)
