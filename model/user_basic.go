package model

import (
	"errors"
	"fmt"
	"star_game/utils"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model

	Uid            string `json:"uid" gorm:"type:varchar(8);not null;unique"`            // 用户的ID
	Account        string `json:"account" gorm:"type:varchar(16);not null;unique"`       // 用户账号
	InvitationCode string `json:"invitationCode" gorm:"type:varchar(6);not null;unique"` // 邀请码
	Password       string `json:"password" gorm:"type:varchar(128);not null"`            // 用户的密码
	Salt           string `json:"salt" gorm:"type:varchar(8);not null:"`                 // 盐
	Avatar         string `json:"avatar" gorm:"type:varchar(128);null"`                  // 用户头像
	Sex            string `json:"sex" gorm:"type:varchar(16);null"`                      // 用户性别
	Dress          string `json:"dress" gorm:"type:varchar(255);null"`                   // 用户地址
	NickName       string `json:"nickName" gorm:"type:varchar(7);null"`                  // 用户的昵称
	Phone          string `json:"phone" gorm:"type:varchar(13);null"`                    // 用户的手机
	Email          string `json:"email" gorm:"type:varchar(128);null"`                   // 用户的邮箱
	SuperiorID     string `json:"superiorId" gorm:"type:varchar(8);null"`                // 上级ID
	CreatIp        string `json:"creatIp" gorm:"type:varchar(64);null"`                  // 注册时的IP地址
	Token          string `json:"token" gorm:"type:varchar(255);null"`                   // Token

	// 邀请码
	// 上级ID
	// 钻石余额
	// 金币余额
	// 佣金
	// 个人的设置选项
}

func (UserBasic) TableName() string {
	return "user_basic"
}

// 添加一个用户
func CreateUser(userBasic *UserBasic) error {
	fmt.Println("=> 正在添加数据", userBasic)

	if result := DB.Create(&userBasic); result.Error != nil {
		fmt.Println("=> 创建用户失败")
		return result.Error
	} else {
		fmt.Println("=> 创建用户成功")
		return nil
	}

}

// 添加内置账号
func CreateBoss() {
	fmt.Println("=> 正在添加内置账号信息...")
	userBasic := UserBasic{}
	userBasic.Account = viper.GetString("boss.account")
	userBasic.Salt = utils.GetCode(8)
	userBasic.InvitationCode = utils.GetCode((6))
	bossPassword := utils.Md5Encode(viper.GetString("boss.password") + viper.GetString("salt.client"))
	userBasic.Password = utils.Crypto(bossPassword, userBasic.Salt)
	userBasic.Uid = viper.GetString("boss.uid")
	userBasic.NickName = viper.GetString("boss.nickName")
	if result := DB.Create(&userBasic); result.Error != nil {
		fmt.Println("=> 创建内置用户失败...")
	} else {
		fmt.Println("=> 创建内置用户成功...")
	}
}

// 查找账号
func FindAccountInUserBasic(account string) (*UserBasic, error) {
	fmt.Println("=> 正在查找账号信息:", account)
	userBasic := UserBasic{}
	if result := DB.First(&userBasic, "account = ?", account); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("=> 查询无结果", result.Error)
			return nil, nil
		}
		fmt.Println("=> 查询有错误", result.Error)
		return nil, result.Error
	} else {
		fmt.Println("=> 查询到数据", userBasic)
		return &userBasic, nil
	}

}

func FindUidByInvitationCode(invitationCode string) (*UserBasic, error) {
	fmt.Println("=> 正在查找邀请码信息:", invitationCode)
	userBasic := UserBasic{}
	if result := DB.First(&userBasic, "invitation_code = ?", invitationCode); result.Error != nil {
		fmt.Println("=> 查询有错误", result.Error)
		return nil, result.Error
	} else {
		fmt.Println("=> 查询到数据", userBasic)
		return &userBasic, nil
	}

}

// 创建数据库
// 如果数据库不存在则创建数据库
func CreateUserBasic() {
	userBasic := UserBasic{}
	if !DB.Migrator().HasTable(userBasic.TableName()) {
		if err := DB.AutoMigrate(&userBasic); err != nil {
			fmt.Println("=> 数据创建失败: ", err)
		} else {
			fmt.Println("=> 数据库创建成功...")
			fmt.Println("=> 正在创建内置账号...")
			CreateBoss()
		}
	}
}

// 更新Token
func UpdataToken(user *UserBasic, value string) error {
	if result := DB.Model(&user).Update("Token", value); result.Error != nil {
		fmt.Println("=> 更新数据库错误: ", result.Error)
		return result.Error
	} else {
		fmt.Println("=> 更新数据库成功")
		return nil
	}

}

// func GetUserList() *gorm.DB {
// 	data := [...]UserBasic{}
// 	resault := DB.Find(&data)
// 	if resault.Error != nil {
// 		fmt.Println("=> error: ", resault.Error)
// 	}
// 	return resault
// }
