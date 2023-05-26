package res

import (
	"fmt"
)

type Room struct {
	Game      string   `json:"game"`
	RoomId    string   `json:"roomId"`
	CreateUid string   `json:"createUid"`
	Current   int      `json:"current"`
	Round     int      `json:"round"`
	GameState string   `json:"gameState"`
	IsAllDrop bool     `json:"isAllDrop"`
	Players   []Player `json:"players"`
	Message   string   `json:"message"`
}

func (romm *Room) SendMessage(data []byte) error {
	for _, player := range romm.Players {
		fmt.Println("=> 正在发消息", player)
		if err := player.Conn.WriteMessage(data); err != nil {
			fmt.Println("=> 对方不在线", player)
			continue
		}
	}
	return nil
}

func (romm *Room) FindUid(uid string) bool {
	for _, player := range romm.Players {
		if player.Uid == uid {
			return true
		}
	}
	return false
}

func (romm *Room) Upgrade(player *Player) {
	for i, _player := range romm.Players {
		if _player.Uid == player.Uid {
			romm.Players[i].Avatar = player.Avatar
			romm.Players[i].Sex = player.Sex
			romm.Players[i].Conn = player.Conn
			romm.Players[i].Fraction = player.Fraction
			romm.Players[i].IsReady = player.IsReady
			romm.Players[i].IsFolded = player.IsFolded
			romm.Players[i].NickName = player.NickName
			romm.Players[i].Uid = player.Uid
		}
	}
}
