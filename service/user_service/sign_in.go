package user_service

import (
	"poker_game/model"
	"poker_game/res"
	"poker_game/utils"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
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
	findAccountResult, findAccountError := model.FindAccount(account)
	if findAccountResult == nil && findAccountError == nil {
		response := res.Response{
			Code:    402,
			Message: "账号不存在",
		}
		response.Send(c)
		return
	}

	if findAccountError != nil {
		response := res.Response{
			Code:    -1,
			Message: "数据库连接失败,请稍后重试",
		}
		response.Send(c)
		return
	}

	if findAccountResult.Password == utils.Crypto(password, findAccountResult.Salt) {
		response := res.Response{
			Code:    200,
			Message: "登陆成功",
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
