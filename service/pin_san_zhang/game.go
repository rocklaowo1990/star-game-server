package pin_san_zhang

import (
	"star_game/common"
	"star_game/res"
)

func GameManagerHanler(room *res.Room, wsReceive *res.WsReceive, conn *common.Connection) error {
	player := res.Player{}
	player.Avatar = wsReceive.Data["avatar"].(string)
	player.Sex = wsReceive.Data["sex"].(string)
	player.NickName = wsReceive.Data["nickName"].(string)
	player.Uid = wsReceive.Data["uid"].(string)
	player.Conn = conn

	switch wsReceive.Type {
	case "enterRoom":
		if err := enterRoom(room, &player); err != nil {
			return err
		}
	case "ready":
		if err := ready(room, &player); err != nil {
			return err
		}
	case "unReady":
		if err := unready(room, &player); err != nil {
			return err
		}
	}

	return nil
}
