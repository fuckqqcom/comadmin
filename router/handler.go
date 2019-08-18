package router

import (
	"comadmin/http/crontroller/admin"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*admin.HttpHandler
}
