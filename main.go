package main

import (
	"poker_game/config"
	"poker_game/model"
	"poker_game/router"
	"poker_game/service"
)

func main() {
	// 初始化配置文件
	config.InitConfig()

	// 初始化 mysql 数据库
	model.InitDB()

	// 初始化 redis
	service.InitRedis()

	go service.Manager.Client()

	// 启动接口监听
	r := router.NewRouter()
	rErr := r.Run(":8080")
	if rErr != nil {
		panic(rErr)
	}

}
