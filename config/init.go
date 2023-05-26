package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("=> 配置信息初始化失败...")
		fmt.Println("=> ", err)
	} else {
		fmt.Println("=> 配置信息初始化成功...")
	}
}
