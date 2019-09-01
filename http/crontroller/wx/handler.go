package wx

import (
	"comadmin/logic/wx"
)

type HttpWxHandler struct {
	//todo log
	logic wx.LogicHandler
}

func NewWxHttpAdminHandler(path string) *HttpWxHandler {
	return &HttpWxHandler{logic: wx.NewLogic(path)}
}
