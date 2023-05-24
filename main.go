package main

import (
	"star_game/config"
	"star_game/model"
	"star_game/router"
	"star_game/service"
)

func main() {
	// 初始化配置文件
	config.InitConfig()

	// 初始化 mysql 数据库
	model.InitDB()

	// 初始化 redis
	service.InitRedis()

	// go service.Manager.Client()

	// 启动接口监听
	r := router.NewRouter()
	rErr := r.Run(":8080")
	if rErr != nil {
		panic(rErr)
	}

}
