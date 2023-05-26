package model

import (
	"fmt"

	"gorm.io/gorm"
)

type RoomBasic struct {
	gorm.Model

	RoomId    string `json:"roomId" gorm:"type:varchar(6);not null;unique"`    // 房间ID
	CreateUid string `json:"createUid" gorm:"type:varchar(8);not null;unique"` // 创建ID
	Current   int    `json:"current" gorm:"type:int;not null;unique"`          // 当前局数
	IsAllDrop int    `json:"isAllDrop" gorm:"int;not null"`                    // 是否全票解散
	Round     int    `json:"round" gorm:"type:int;not null:"`                  // 总局数
	GameName  string `json:"gameName" gorm:"type:varchar(128);null"`           // 游戏名臣恶搞
	GameState string `json:"gameState" gorm:"type:varchar(16);null"`           // 游戏状态
	Players   string `json:"dress" gorm:"type:varchar(255);null"`              // 玩家列表
}

func (RoomBasic) TableName() string {
	return "room_basic"
}

// 创建数据库
// 如果数据库不存在则创建数据库
func CreateRoomBasic() {
	roomBasic := RoomBasic{}
	if !DB.Migrator().HasTable(roomBasic.TableName()) {
		if err := DB.AutoMigrate(&roomBasic); err != nil {
			fmt.Println("=> 数据创建失败: ", err)
			panic(err)
		}
		fmt.Println("=> 数据库创建成功...")
	}
}
