package app

import (
	"github.com/gin-gonic/gin"
	"wxadmin/pkg/e"
)

type GContext = *gin.Context

type G struct {
	*gin.Context
}

func (g *G) Json(httpCode, code int, action, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = e.GetMsg(code)
	m["action"] = action
	m["data"] = data
	g.JSON(httpCode, m)
	return
}
