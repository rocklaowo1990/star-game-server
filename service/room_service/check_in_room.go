package room_service

import (
	"star_game/res"

	"github.com/gin-gonic/gin"
)

func CheckInRoom(c *gin.Context) {
	uid := c.Query("uid")

	if room := roomManager.CheckInRoom(uid); room != nil {
		data := make(map[string]interface{})
		data["roomId"] = room.RoomId
		data["game"] = room.Game

		response := res.Response{
			Code:    200,
			Message: "获取房间信息成功",
			Data:    data,
		}

		response.Send(c)
		return
	}

	response := res.Response{
		Code:    -1,
		Message: "房间信息初始化成功",
	}
	response.Send(c)
}
