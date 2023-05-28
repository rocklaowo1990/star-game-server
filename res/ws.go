package res

// ws 接收到的数据格式
type WsReceive struct {
	Game   string         `json:"game"`
	Type   string         `json:"type"`
	RoomId string         `json:"roomId"`
	Data   map[string]any `json:"data"`
}

// ws 回复的格式
type WsResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
