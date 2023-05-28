package pin_san_zhang

import (
	"encoding/json"
	"fmt"
	"star_game/res"
)

func ready(room *res.Room, player *res.Player) error {
	room.Message = fmt.Sprintf("%s 已准备", player.NickName)

	wsResponse := res.WsResponse{}
	wsResponse.Type = "game"
	wsResponse.Data = room

	room.UpdateIsReady(player.Uid, true)

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

func unready(room *res.Room, player *res.Player) error {
	room.Message = fmt.Sprintf("%s 已取消准备", player.NickName)

	wsResponse := res.WsResponse{}
	wsResponse.Type = "game"
	wsResponse.Data = room

	room.UpdateIsReady(player.Uid, false)

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
