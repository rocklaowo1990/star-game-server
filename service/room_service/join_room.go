package room_service

import (
	"star_game/res"

	"github.com/gin-gonic/gin"
)

func JoinRoom(c *gin.Context) {
	roomId := c.Query("roomId")
	uid := c.Query("uid")

	if isFindRoomId := roomManager.FindRoomId(roomId); !isFindRoomId {
		response := res.Response{
			Code:    402,
			Message: "房间不存在",
		}
		response.Send(c)
		return
	}

	if isFindUid := roomManager.FindUid(roomId, uid); isFindUid {
		response := res.Response{
			Code:    403,
			Message: "你已加入其他房间",
		}
		response.Send(c)
		return
	}

	data := make(map[string]interface{})
	data["roomId"] = roomId

	response := res.Response{
		Code:    200,
		Data:    data,
		Message: "加入房间成功",
	}
	response.Send(c)
}
