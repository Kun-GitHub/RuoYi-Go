// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Author: K.
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

package rywebsocket

import (
	"RuoYi-Go/pkg/logger"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"go.uber.org/zap"
)

var webSocket *iris.Application

func InitWebSocket(w *iris.Application) {
	webSocket = w
}

func StartWebSocket() {
	webSocket.Get("/ws", websocket.Handler(initWebsocket()))
}

// InitConfig 函数中使用viper读取配置文件并映射到AppConfig结构体
func initWebsocket() *neffos.Server {
	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			logger.Log.Info(fmt.Sprintf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID()))

			//mg := websocket.Message{
			//	Body:     msg.Body,
			//	IsNative: true,
			//}
			//
			//nsConn.Conn.Write(mg)
			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		logger.Log.Info(fmt.Sprintf("[%s] Connected to server!", c.ID()))
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		logger.Log.Info(fmt.Sprintf("[%s] Disconnected from server", c.ID()))
	}

	ws.OnUpgradeError = func(err error) {
		logger.Log.Error("Upgrade Error: %v", zap.Error(err))
	}

	return ws
}
