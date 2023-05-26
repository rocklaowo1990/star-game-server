package user_service

import (
	"fmt"
	"star_game/model"
	"star_game/res"
	"star_game/utils"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Uid            string `json:"uid"`            // 用户的ID
	Account        string `json:"account"`        // 用户账号
	InvitationCode string `json:"invitationCode"` // 邀请码
	Avatar         string `json:"avatar"`         // 用户头像
	Sex            string `json:"sex"`            // 用户性别
	Dress          string `json:"dress"`          // 用户地址
	NickName       string `json:"nickName"`       // 用户的昵称
	Phone          string `json:"phone"`          // 用户的手机
	Email          string `json:"email"`          // 用户的邮箱
	SuperiorID     string `json:"superiorID"`     // 上级ID
	CreatIp        string `json:"creatIp"`        // 注册时的IP地址
	Token          string `json:"token"`          // token
	RoomId         string `json:"roomId"`         // 房间ID
	Gold           int64  `json:"gold"`           // 金币
}

func SignIn(c *gin.Context) {
	var (
		result *model.UserBasic
		err    error
	)
	// 读取客户端传过来的数据
	account := c.Query("account")
	password := c.Query("password")

	// 如果客户端传过来的用户名或密码为空
	if account == "" || password == "" {
		response := res.Response{
			Code:    401,
			Message: "用户名或密码不能为空",
		}
		response.Send(c)
		return
	}

	// 查找数据库是否存在相同的账号信息
	if result, err = model.FindAccountInUserBasic(account); err == nil && result == nil {
		response := res.Response{
			Code:    402,
			Message: "账号不存在",
		}
		response.Send(c)
		return
	}

	if err != nil {
		response := res.Response{
			Code:    -1,
			Message: "数据库连接失败,请稍后重试",
		}
		response.Send(c)
		return
	}

	if result.Password == utils.Crypto(password, result.Salt) {
		var token string
		if token, err = utils.GenerateToken(result.Uid); err != nil {
			fmt.Println("=> 创建 token 出错", err)
			return
		}

		if err = model.UpdataToken(result, token); err != nil {
			fmt.Println("=> 更新 token 出错", err)
			return
		}

		data := UserInfo{
			Uid:            result.Uid,
			Account:        result.Account,
			InvitationCode: result.InvitationCode,
			Avatar:         result.Avatar,
			Sex:            result.Sex,
			Dress:          result.Dress,
			NickName:       result.NickName,
			Phone:          result.Phone,
			Email:          result.Email,
			SuperiorID:     result.SuperiorID,
			CreatIp:        result.CreatIp,
			Token:          token,
		}

		response := res.Response{
			Code:    200,
			Message: "登陆成功",
			Data:    data,
		}

		response.Send(c)

		return
	} else {
		response := res.Response{
			Code:    403,
			Message: "用户名或密码错误",
		}
		response.Send(c)
		return
	}

}
