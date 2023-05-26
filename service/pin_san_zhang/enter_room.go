package pin_san_zhang

import (
	"encoding/json"
	"fmt"
	"star_game/common"
	"star_game/res"
)

func enterRoom(room *res.Room, gameMessage *res.GameMessage, conn *common.Connection) error {
	player := res.Player{}
	player.Avatar = gameMessage.Data["avatar"].(string)
	player.Sex = gameMessage.Data["sex"].(string)
	player.Conn = conn
	player.Fraction = 0
	player.IsReady = false
	player.IsFolded = false
	player.NickName = gameMessage.Data["nickName"].(string)
	player.Uid = gameMessage.Data["uid"].(string)

	if isFindUidFromRoom := room.FindUid(player.Uid); !isFindUidFromRoom {
		room.Players = append(room.Players, player)
		room.Message = fmt.Sprintf("%s 加入房间", player.NickName)
	} else {
		room.Message = fmt.Sprintf("%s 返回房间", player.NickName)
		room.Upgrade(&player)
	}

	var data []byte
	var err error

	if data, err = json.Marshal(&room); err != nil {
		return err
	}

	if err := room.SendMessage(data); err != nil {
		return err
	}

	return nil
}
