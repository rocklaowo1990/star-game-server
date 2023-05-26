package main

import (
	"star_game/common"
	"star_game/config"
	"star_game/model"
	"star_game/router"
)

func main() {
	// 初始化配置文件
	config.InitConfig()

	// 初始化 mysql 数据库
	model.InitDB()

	// 初始化 common
	common.InitCommon()

	// go service.Manager.Client()

	// 启动接口监听
	r := router.NewRouter()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
