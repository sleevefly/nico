package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 20 * time.Second,
	// 取消 ws 跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(c *gin.Context) {
	// 获取WebSocket连接
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorf("upgrade contral err %s", err.Error())
		return
	}
	// 处理WebSocket消息
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println("messageType:", messageType)
		fmt.Println("p:", string(p))

		// 输出WebSocket消息内容
		c.Writer.Write(p)
	}

	// 关闭WebSocket连接
	ws.Close()
}
