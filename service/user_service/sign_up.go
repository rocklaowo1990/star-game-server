package user_service

import (
	"fmt"
	"star_game/model"
	"star_game/res"
	"star_game/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SignUp(c *gin.Context) {
	user := new(model.UserBasic)

	// 读取客户端传过来的数据
	account := c.Query("account")
	password := c.Query("password")
	invitationCode := c.Query("invitationCode")

	fmt.Println(invitationCode)

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
	// 查询失败的情况
	findAccountResult, findAccountError := model.FindAccountInUserBasic(account)
	if findAccountError != nil {
		response := res.Response{
			Code:    -1,
			Message: "数据库连接失败,请稍后重试",
		}
		response.Send(c)
		return
	}

	// 查询到有数据的情况
	if findAccountResult != nil {
		response := res.Response{
			Code:    402,
			Message: "用户已被注册",
		}
		response.Send(c)
		return
	}

	// 开始为注册的用户初始化信息
	// 注册IP
	if c.ClientIP() == "::1" {
		user.CreatIp = "127.0.0.1"
	} else {
		user.CreatIp = c.ClientIP()
	}

	// 上级ID
	if invitationCode == "" {
		user.SuperiorID = viper.GetString("boss.uid")
	} else {
		// 如果用户填写了邀请码
		// 需要先去数据查到拥有该邀请码的用户UID
		// 然后将查到到的UID设定为注册用户的上级ID
		findUidByInvitationCodeResult, findUidByInvitationCodeError := model.FindUidByInvitationCode(invitationCode)
		if findUidByInvitationCodeError != nil {
			response := res.Response{
				Code:    403,
				Message: "邀请码输入有误",
			}
			response.Send(c)
			return
		}

		if findUidByInvitationCodeResult != nil {
			user.SuperiorID = findUidByInvitationCodeResult.Uid
		}
	}

	// 给密码加密
	user.Salt = utils.GetCode(8)
	user.Password = utils.Crypto(password, user.Salt)
	// UID
	user.Uid = utils.GetID(8, false)
	// 昵称
	user.NickName = fmt.Sprintf("玩家%s", utils.GetID(5, true))
	// 为用户生成邀请码
	user.InvitationCode = utils.GetCode(6)
	// 账号为用户传过来的账号
	user.Account = account

	// 开始将用户信息写入数据库
	createUserError := model.CreateUser(user)
	// 写入数据库失败的情况
	if createUserError != nil {
		response := res.Response{
			Code:    404,
			Message: "写入数据失败, 请稍后重试",
		}
		response.Send(c)
		return
	} else {
		// 写入成功并提示用户注册成功
		response := res.Response{
			Code:    200,
			Message: "注册成功",
		}
		response.Send(c)
		return
	}

}
