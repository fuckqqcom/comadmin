package router

import (
	"comadmin/http/crontroller"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	r *gin.Engine
	h *crontroller.HttpHandler
}
