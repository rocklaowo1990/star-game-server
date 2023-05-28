package res

import (
	"encoding/json"
	"star_game/common"

	"github.com/robfig/cron/v3"
)

type Room struct {
	Game      string     `json:"game"`
	RoomId    string     `json:"roomId"`
	CreateUid string     `json:"createUid"`
	Current   int        `json:"current"`
	Round     int        `json:"round"`
	GameState string     `json:"gameState"`
	IsAllDrop bool       `json:"isAllDrop"`
	Players   []Player   `json:"players"`
	Message   string     `json:"message"`
	Timer     *cron.Cron `json:"times"`
}

func (room *Room) SendMessage(data []byte) error {
	for _, player := range room.Players {
		if err := player.Conn.WriteMessage(data); err != nil {
			continue
		}
	}
	return nil
}

func (room *Room) TimerStart() error {
	wsResponse := WsResponse{}
	wsResponse.Type = "timer"

	i := 0
	room.Timer = cron.New(cron.WithSeconds())
	spec := "*/1 * * * * ?"
	if _, err := room.Timer.AddFunc(spec, func() {
		i++

		data := make(map[string]interface{})
		data["time"] = i

		wsResponse.Data = data

		var _data []byte

		_data, _ = json.Marshal(&wsResponse)

		for _, player := range room.Players {
			if err := player.Conn.WriteMessage(_data); err != nil {
				continue
			}
		}
	}); err != nil {
		return err
	}

	room.Timer.Start()
	return nil
}

func (room *Room) FindUid(uid string) bool {
	for _, player := range room.Players {
		if player.Uid == uid {
			return true
		}
	}
	return false
}

func (room *Room) UpdateConn(uid string, conn *common.Connection) {
	for i, _player := range room.Players {
		if _player.Uid == uid {
			room.Players[i].Conn = conn
		}
	}
}

func (room *Room) UpdateIsReady(uid string, result bool) {
	for i, _player := range room.Players {
		if _player.Uid == uid {
			room.Players[i].IsReady = result
		}
	}
}
