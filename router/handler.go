package router

import (
	"comadmin/http/crontroller/admin"
	"comadmin/http/crontroller/wx"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*admin.HttpAdminHandler
	*wx.HttpWxHandler
}
