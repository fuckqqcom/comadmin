package admin

import "comadmin/logic/admin"

type HttpHandler struct {
	logic admin.LogicHandler
}

func NewAdminHttpHandler(path string) *HttpHandler {
	return &HttpHandler{logic: admin.NewLogic(path)}
}
