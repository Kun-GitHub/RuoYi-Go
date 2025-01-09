// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package ryws

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"go.uber.org/zap"
)

type WebsocketServer struct {
	logger *zap.Logger
	conns  map[string]*websocket.Conn
}

func StartWebSocket(app *iris.Application, l *zap.Logger) *WebsocketServer {
	conns := make(map[string]*websocket.Conn)
	ws := &WebsocketServer{
		logger: l,
		conns:  conns,
	}

	ws.logger.Info("websocket start")
	app.Get("/websocket", websocket.Handler(ws.handlerWebsocket()))
	return ws
}

func (this *WebsocketServer) handlerWebsocket() *neffos.Server {
	ws := websocket.New(websocket.DefaultGobwasUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			this.logger.Info(fmt.Sprintf("Server got: %s from [%s]\n", msg.Body, nsConn.Conn.ID()))

			//mg := websocket.Message{
			//	Body:     []byte(jsonBytes),
			//	IsNative: true,
			//}
			//c.Write(mg)
			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		this.logger.Info(fmt.Sprintf("[%s] Connected to server!\n", c.ID()))
		this.conns[c.ID()] = c

		//mg := websocket.Message{
		//	Body:     []byte(jsonBytes),
		//	IsNative: true,
		//}
		//c.Write(mg)
		return nil
	}
	ws.OnDisconnect = func(c *websocket.Conn) {
		delete(this.conns, c.ID())
		this.logger.Info(fmt.Sprintf("[%s] Disconnected from server\n", c.ID()))
	}
	ws.OnUpgradeError = func(err error) {
		this.logger.Error("Upgrade Error: %v\n", zap.Error(err))
	}

	return ws
}
