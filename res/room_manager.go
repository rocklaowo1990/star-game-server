package res

import (
	"encoding/json"
	"fmt"
	"star_game/common"
)

type RoomManager struct {
	Rooms []Room `json:"rooms"`
}

func (roomManager *RoomManager) FindRoomId(roomId string) bool {
	for _, room := range roomManager.Rooms {
		if room.RoomId == roomId {
			return true
		}
	}
	return false
}

func (roomManager *RoomManager) FindUid(roomId string, uid string) bool {
	for _, room := range roomManager.Rooms {
		if room.RoomId == roomId {
			continue
		}

		for _, player := range room.Players {
			if player.Uid == uid {
				return true
			}
		}
	}
	return false
}

func (roomManager *RoomManager) CheckCreateRoomMax(uid string) int {
	count := 0
	for _, room := range roomManager.Rooms {
		if room.CreateUid == uid {
			count++
		}
	}
	return count
}

func (roomManager *RoomManager) CheckInRoom(uid string) *Room {
	for _, room := range roomManager.Rooms {
		for _, player := range room.Players {
			if player.Uid == uid {
				return &room
			}
		}
	}
	return nil
}

func (roomManager *RoomManager) OutHandler(conn *common.Connection) error {
	for i, room := range roomManager.Rooms {
		for _, player := range room.Players {
			if player.Conn == conn {
				message := fmt.Sprintf("%s 断开连接", player.NickName)
				roomManager.Rooms[i].Message = message

				var data []byte
				var err error

				if data, err = json.Marshal(roomManager.Rooms[i]); err != nil {
					return err
				}

				roomManager.Rooms[i].SendMessage(data)
				return nil
			}
		}
	}
	return nil
}
