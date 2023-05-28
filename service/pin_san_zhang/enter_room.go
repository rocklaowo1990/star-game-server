package pin_san_zhang

import (
	"encoding/json"
	"fmt"
	"star_game/res"
)

func enterRoom(room *res.Room, player *res.Player) error {
	player.Fraction = 0
	player.IsReady = false
	player.IsFolded = false

	if len(room.Players) == 0 {
		room.TimerStart()
	}

	wsResponse := res.WsResponse{}
	wsResponse.Type = "game"

	if isFindUidFromRoom := room.FindUid(player.Uid); !isFindUidFromRoom {
		room.Players = append(room.Players, *player)
		room.Message = fmt.Sprintf("%s 加入房间", player.NickName)
	} else {
		room.UpdateConn(player.Uid, player.Conn)
		room.Message = fmt.Sprintf("%s 返回房间", player.NickName)
	}

	wsResponse.Data = room

	var data []byte
	var err error

	if data, err = json.Marshal(&wsResponse); err != nil {
		return err
	}

	if err := room.SendMessage(data); err != nil {
		return err
	}

	return nil
}
