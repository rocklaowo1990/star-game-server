package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

func (r *Response) Test() {
	fmt.Println(r)
}

// func main() {
// 	u := Response{
// 		Code:    200,
// 		Message: "123123123",
// 	}
// 	u.Test()
// }

func main() {
	i := 0
	c := cron.New(cron.WithSeconds())
	spec := "*/2 * * * * ?"
	_, err := c.AddFunc(spec, func() {
		i++
		fmt.Println("cron times : ", i)
	})
	if err != nil {
		fmt.Println("AddFunc error : ", err)
		return
	}
	c.Start()

	defer c.Stop()
	select {}
}
