package res

type GameMessage struct {
	Game   string         `json:"game"`
	Type   string         `json:"type"`
	RoomId string         `json:"roomId"`
	Data   map[string]any `json:"data"`
}
