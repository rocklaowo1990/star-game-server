package res

import "star_game/common"

type Player struct {
	Conn     *common.Connection `json:"conn"`
	Uid      string             `json:"uid"`
	NickName string             `json:"nickName"`
	Avatar   string             `json:"avatar"`
	IsReady  bool               `json:"isReady"`
	IsFolded bool               `json:"isFolded"`
	Sex      string             `json:"sex"`
	Fraction int                `json:"fraction"`
}
