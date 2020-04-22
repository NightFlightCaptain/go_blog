package pkg

import "go_blog/pkg/e"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetResponse(code int, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	}
}
