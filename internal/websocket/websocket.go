package ws

import (
	"fmt"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"go.uber.org/zap"
	"time-machine/pkg/logger"
)

// InitConfig 函数中使用viper读取配置文件并映射到AppConfig结构体
func InitWebsocket() *neffos.Server {
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
