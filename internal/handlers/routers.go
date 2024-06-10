// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package handler

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/middlewares"
	"RuoYi-Go/internal/responses"
	"RuoYi-Go/internal/services"
	"RuoYi-Go/pkg/logger"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"net/url"
	"strings"
)

func GetRouters(ctx iris.Context) {
	loginUser := middlewares.GetLoginUser()
	if loginUser == nil || loginUser.UserID == 0 {
		ctx.JSON(responses.Error(iris.StatusUnauthorized, "请重新登录"))
		return
	}

	var menus = make([]*services.SysMenuStruct, 0)
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
	ctx.JSON(responses.Success(buildMenuTree(menus)))
}

type MetaStruct struct {
	Title   string `json:"title,omitempty"`
	Icon    string `json:"icon,omitempty"`
	NoCache bool   `json:"noCache,omitempty"`
	Link    string `json:"link,omitempty"`
}

type routerStruct struct {
	Hidden     bool            `json:"hidden,omitempty"`
	AlwaysShow bool            `json:"alwaysShow,omitempty"`
	Name       string          `json:"name,omitempty"`
	Path       string          `json:"path,omitempty"`
	Component  string          `json:"component,omitempty"`
	Redirect   string          `json:"redirect,omitempty"`
	Query      string          `json:"query,omitempty"`
	Meta       *MetaStruct     `json:"meta,omitempty"`
	Children   []*routerStruct `json:"children,omitempty"`
}

// BuildMenuTree builds the menu tree from a flat list of SysMenu.
func buildMenuTree(menus []*services.SysMenuStruct) []*routerStruct {
	menuMap := make(map[int64]*services.SysMenuStruct)
	rootMenus := make([]*routerStruct, 0)

	// Fill the map with all menus
	for _, menu := range menus {
		menuMap[menu.MenuID] = menu
	}

	// Construct the tree
	for _, menu := range menus {
		if menu.ParentID == 0 {
			router := &routerStruct{
				Hidden:    menu.Visible == "1",
				Name:      getRouteName(menu),
				Path:      getRouterPath(menu),
				Component: getComponent(menu),
				Query:     menu.Query,
				Meta: &MetaStruct{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: menu.IsCache == 1,
					Link:    menu.Path,
				},
			}

			rootMenus = append(rootMenus, router)
		} else {
			parent, exists := menuMap[menu.ParentID]
			if exists {
				//parent.Children = append(parent.Children, menu)
				// Concatenate the path with parent's path
				menu.Path = parent.Path + "/" + menu.Path
			}
		}
	}

	return rootMenus
}

//func buildMenus(menus []*services.SysMenuStruct) []*routerStruct {
//	var routers []*routerStruct
//	for _, menu := range menus {
//		router := &routerStruct{
//			Hidden:    strings.EqualFold(menu.Visible, "1"),
//			Name:      getRouteName(menu),
//			Path:      getRouterPath(menu),
//			Component: getComponent(menu),
//			Query:     menu.Query,
//			Meta: &MetaStruct{
//				Title:   menu.MenuName,
//				Icon:    menu.Icon,
//				NoCache: menu.IsCache == 1,
//				Link:    menu.Path,
//			},
//		}
//
//		children := menu.Children
//		if len(children) > 0 && menu.MenuType == "TYPE_DIR" {
//			router.AlwaysShow = true
//			router.Redirect = "noRedirect"
//			router.Children = buildMenus(children)
//		} else if menu.ParentID == 0 && isInnerLink(menu) {
//			router.Meta = &MetaStruct{Title: menu.MenuName, Icon: menu.Icon}
//			router.Path = "/"
//			childrenList := make([]*routerStruct, 1)
//			childrenList[0] = &routerStruct{
//				Path:      innerLinkReplaceEach(menu.Path),
//				Component: "INNER_LINK",
//				Name:      strings.Title(strings.Replace(menu.Path, "-", " ", -1)),
//				Meta: &MetaStruct{
//					Title: menu.MenuName,
//					Icon:  menu.Icon,
//					Link:  menu.Path,
//				},
//			}
//			router.Children = childrenList
//		} else if isMenuFrame(menu) {
//			router.Meta = nil
//			childrenList := make([]*routerStruct, 1)
//			childrenList[0] = &routerStruct{
//				Path:      menu.Path,
//				Component: menu.Component,
//				Name:      strings.Title(strings.Replace(menu.Path, "-", " ", -1)),
//				Meta: &MetaStruct{
//					Title:   menu.MenuName,
//					Icon:    menu.Icon,
//					NoCache: menu.IsCache == 1,
//					Link:    menu.Path,
//				},
//				Query: menu.Query,
//			}
//			router.Children = childrenList
//		}
//		routers = append(routers, router)
//	}
//	return routers
//}

// 注意：以下辅助函数需要根据实际情况实现
func getRouteName(menu *services.SysMenuStruct) string {
	routerName := strings.Title(menu.Path)
	// Non-outer link and is a first-level directory (type is directory)
	if isMenuFrame(menu) {
		routerName = ""
	}
	return routerName
}

func getRouterPath(menu *services.SysMenuStruct) string {
	routerPath := menu.Path

	// Inner link open external way
	if menu.ParentID != 0 && isInnerLink(menu) {
		routerPath = innerLinkReplaceEach(routerPath)
	}

	// Not an outer link and is a top-level directory (type is directory)
	if menu.ParentID == 0 && menu.MenuType == "TYPE_DIR" && menu.IsFrame == common.NO_FRAME {
		routerPath = "/" + menu.Path
	} else if isMenuFrame(menu) {
		routerPath = "/"
	}

	return routerPath
}

func getComponent(menu *services.SysMenuStruct) string {
	component := "LAYOUT"
	if strings.TrimSpace(menu.Component) != "" && !isMenuFrame(menu) {
		component = menu.Component
	} else if strings.TrimSpace(menu.Component) == "" && menu.ParentID != 0 && isInnerLink(menu) {
		component = "INNER_LINK"
	} else if strings.TrimSpace(menu.Component) == "" && isParentView(menu) {
		component = "PARENT_VIEW"
	}
	return component
}

func isMenuFrame(menu *services.SysMenuStruct) bool {
	return menu.ParentID == 0 && common.TYPE_MENU == menu.MenuType && menu.IsFrame == common.NO_FRAME
}

func isInnerLink(menu *services.SysMenuStruct) bool {
	return menu.IsFrame == common.NO_FRAME && isHTTP(menu.Path)
}

// isParentView checks if the given menu is a parent view.
// This function needs to be implemented based on your specific logic.
func isParentView(menu *services.SysMenuStruct) bool {
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
	return replaceEach(path)
}

const (
	HTTP  = "http://"
	HTTPS = "https://"
	WWW   = "www."
	DOT   = "."
	COLON = ":"
	SLASH = "/"
)

func replaceEach(path string) string {
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
