package res

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	// Error   string      `json:"error"`
}

func (response *Response) Send(c *gin.Context) {
	fmt.Println("=> 正在给客户端发送消息", response)
	c.JSON(200, response)
}
