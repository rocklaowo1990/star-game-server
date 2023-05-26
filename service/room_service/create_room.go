package room_service

import (
	"star_game/res"
	"star_game/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	room := res.Room{}
	room.Game = c.Query("game")
	room.CreateUid = c.Query("createUid")
	room.RoomId = utils.GetID(6, false)
	room.Current = 0
	room.IsAllDrop = c.GetBool("isAllDrop")
	room.GameState = "ready"

	var round int
	var isAllDrop bool
	var err error

	if isAllDrop, err = strconv.ParseBool(c.Query("isAllDrop")); err != nil {
		response := res.Response{
			Code:    402,
			Message: "传参错误",
		}
		response.Send(c)
		return
	}

	room.IsAllDrop = isAllDrop

	if round, err = strconv.Atoi(c.Query("round")); err != nil {
		response := res.Response{
			Code:    402,
			Message: "传参错误",
		}
		response.Send(c)
		return
	}

	room.Round = round

	if checkCreateId := roomManager.CheckCreateRoomMax(room.CreateUid); checkCreateId >= 3 {
		response := res.Response{
			Code:    401,
			Message: "房间创建达到最大限制",
		}
		response.Send(c)
		return
	}

	if isFindUid := roomManager.FindUid(room.RoomId, room.CreateUid); isFindUid {
		response := res.Response{
			Code:    403,
			Message: "你已加入其他房间",
		}
		response.Send(c)
		return
	}

	for {
		if isFindRoomId := roomManager.FindRoomId(room.RoomId); isFindRoomId {
			room.RoomId = utils.GetID(6, false)
		} else {
			break
		}
	}

	roomManager.Rooms = append(roomManager.Rooms, room)

	data := make(map[string]interface{})
	data["roomId"] = room.RoomId

	response := res.Response{
		Code:    200,
		Message: "房间创建成功",
		Data:    data,
	}
	response.Send(c)
}
