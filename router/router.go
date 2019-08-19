package router

import (
	"comadmin/http/crontroller/admin"
	"comadmin/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := &Engine{gin.New(), admin.NewAdminHttpHandler("config/config.json")}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.JWT())
	domain := r.Group("/v1/domain")
	{
		//域
		domain.POST("/create", r.CreateDomain)
		domain.POST("/delete", r.DeleteDoDomain)
		domain.POST("/find", r.FindDomainArgs)
		domain.POST("/update", r.UpdateDomain)
		//app
		domain.POST("/createApp", r.CreateDomainApp)
		domain.POST("/updateApp", r.UpdateDomainApp)
		domain.POST("/findApp", r.FindDomainApp)
		domain.POST("/deleteApp", r.DeleteDomainApp)
	}

	//用户操作
	user := r.Group("/v1/user")
	user.POST("/register", r.Register) //用户注册
	user.POST("/login", r.Login)       //用户登录
	{
		user.POST("/createAdmin", r.RegisterAdminUser)
		user.POST("/delete", r.DeleteUser)
		user.POST("/find", r.FindUser)
		user.POST("/updatePhone", r.UpdateUserPhone)
		user.POST("/updatePwd", r.UpdateUserPwd)
		user.POST("/forbidUser", r.UpdateUserStatus)
	}
	return r.Engine
}
