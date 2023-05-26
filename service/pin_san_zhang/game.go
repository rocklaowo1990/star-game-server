package pin_san_zhang

import (
	"fmt"
	"star_game/common"
	"star_game/res"
)

func GameManagerHanler(room *res.Room, gameMessage *res.GameMessage, conn *common.Connection) error {
	fmt.Println(gameMessage.Type)
	switch gameMessage.Type {
	case "enterRoom":
		enterRoom(room, gameMessage, conn)
	}
	return nil
}
