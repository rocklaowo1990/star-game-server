package room_service

import (
	"encoding/json"
	"fmt"
	"star_game/common"
	"star_game/res"
	"star_game/service/pin_san_zhang"

	"github.com/gin-gonic/gin"
)

var (
	roomManager *res.RoomManager = new(res.RoomManager)
)

func RoomManagerHandler(c *gin.Context, conn *common.Connection) error {
	var (
		data []byte
		err  error
	)

	if data, err = conn.ReadMessage(); err != nil {
		if err = roomManager.OutHandler(conn); err != nil {
			return err
		}
	}

	wsReceive := res.WsReceive{}

	// 解析用户传入的数据
	if err = json.Unmarshal(data, &wsReceive); err != nil {
		fmt.Println("err", err)
		return err
	}

	// 根据传过来的 roomId 寻找到房间 room
	for index, room := range roomManager.Rooms {
		if room.RoomId == wsReceive.RoomId {
			switch wsReceive.Game {
			case "pin_san_zhang":
				if err := pin_san_zhang.GameManagerHanler(&roomManager.Rooms[index], &wsReceive, conn); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
