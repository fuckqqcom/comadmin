package router

import (
	"comadmin/http/crontroller/admin"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := &Engine{r: gin.New(), h: admin.NewAdminHttpHandler("config/config.json")}
	r.r.Use(gin.Logger())
	r.r.Use(gin.Recovery())
	domain := r.r.Group("/domain")
	{
		domain.POST("/create", r.h.CreateDomain)
		domain.GET("/delete", r.h.DeleteDoDomain)
		domain.POST("/find", r.h.FindDomainArgs)
	}
	return r.r
}
