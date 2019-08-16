package app

import (
	"comadmin/pkg/e"
	"github.com/gin-gonic/gin"
)

type GContext = *gin.Context

type G struct {
	*gin.Context
}

func (g *G) Json(httpCode, code int, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = e.GetMsg(code)
	m["data"] = data
	g.JSON(httpCode, m)
}
