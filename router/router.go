package router

import (
	"poker_game/service"
	"poker_game/service/user_service"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/signup", user_service.SignUp)
		user.GET("/signin", user_service.SignIn)
	}

	r.GET("/ws", service.Handler)

	return r
}
