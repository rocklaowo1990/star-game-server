package router

import (
	"net/http"
	"star_game/service"
	"star_game/service/room_service"
	"star_game/service/user_service"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors())

	user := r.Group("/user")
	{
		user.POST("/signup", user_service.SignUp)
		user.GET("/signin", user_service.SignIn)
	}

	room := r.Group("/room")
	{
		room.POST("/create", room_service.CreateRoom)
		room.POST("/join", room_service.JoinRoom)
		room.GET("/checkinroom", room_service.CheckInRoom)
	}

	config := r.Group("/config")
	{
		config.GET("/getconfig", service.GetConfig)
	}

	r.GET("/ws", service.WebSocketHandler)

	return r
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, Access-Control-Allow-Methods")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
