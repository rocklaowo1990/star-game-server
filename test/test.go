package main

import "fmt"

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

func (r *Response) Test() {
	fmt.Println(r)
}

func main() {
	u := Response{
		Code:    200,
		Message: "123123123",
	}
	u.Test()
}
