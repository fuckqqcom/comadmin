package crontroller

import "comadmin/logic/admin"

type HttpHandler struct {
	logic admin.LogicHandler
}

func NewHttpHandler(path string) *HttpHandler {
	return &HttpHandler{logic: admin.NewLogic(path)}
}
