package service

import (
	"fmt"
	"star_game/res"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Salt string `json:"salt"`
}

func GetConfig(c *gin.Context) {
	data := Config{
		Salt: viper.GetString("salt.client"),
	}
	fmt.Println(data)
	response := res.Response{
		Code:    200,
		Message: "获取配置成功",
		Data:    data,
	}
	response.Send(c)
}
