package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Data interface{} `json:"data,omitempty"`
	Flag bool        `json:"flag"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
}

func Response(w http.ResponseWriter, data interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = 400
		body.Msg = err.Error()
		body.Data = data
		body.Flag = false
	} else {
		body.Msg = "OK"
		body.Data = data
		body.Flag = true
	}
	httpx.WriteJson(w, http.StatusBadRequest, body)
}

type code int

const (
	ResCode        code = iota + 199
	ResSuccessCode      // 前端200的自定义状态码 不再渲染，
)
const (
	MsgSuccess = "Success"
)

func NewResponse(flag bool, msg string, data interface{}, code int) (resp *Body) {
	resp = &Body{
		Data: data,
		Flag: flag,
		Msg:  msg,
		Code: code,
	}
	return
}
