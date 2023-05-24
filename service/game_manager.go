package service

import (
	"encoding/json"
	"fmt"
)

type GameMessage struct {
	Game string         `json:"game"`
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

func GameManager(conn *Connection, gameMessage *GameMessage) error {
	if gameMessage.Game == "pin_san_zhang" {
		fmt.Println("=> 正在处理‘拼三张’游戏的业务")
	}

	data, err := json.Marshal("1123123")

	if err != nil {
		return err
	}

	if err = conn.WriteMessage(data); err != nil {
		return err

	}
	return nil
}
