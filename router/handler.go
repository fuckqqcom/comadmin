package router

import (
	"comadmin/http/crontroller/admin"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	r *gin.Engine
	h *admin.HttpHandler
}
