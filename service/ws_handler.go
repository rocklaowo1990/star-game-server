package service

import (
	"net/http"
	"star_game/common"
	"star_game/service/room_service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WebSocketHandler(c *gin.Context) {
	var (
		//websocket 长连接
		ws   *websocket.Conn
		err  error
		conn *common.Connection
	)

	upgrade := websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//header中添加Upgrade:websocket
	if ws, err = upgrade.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}

	if conn, err = common.InitConnection(ws); err != nil {
		goto ERR
	}

	for {
		if err = room_service.RoomManagerHandler(c, conn); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}
