// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package ryws

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func StartWebSocket(ws *iris.Application, l *zap.Logger) {
	logger = l

	logger.Info("websocket start")
	ws.Get("/ws", websocket.Handler(handlerWebsocket()))
}

func handlerWebsocket() *neffos.Server {
	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			logger.Debug(fmt.Sprintf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID()))

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
		logger.Info(fmt.Sprintf("Connected to server! [%s]", c.ID()))
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		logger.Info(fmt.Sprintf("[%s] Disconnected from server", c.ID()))
	}

	ws.OnUpgradeError = func(err error) {
		logger.Error("Upgrade Error: %v", zap.Error(err))
	}

	return ws
}
