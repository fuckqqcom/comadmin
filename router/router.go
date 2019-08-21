package router

import (
	"comadmin/http/crontroller/admin"
	"comadmin/http/crontroller/wx"
	"comadmin/http/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	path := "config/config.json"
	r := &Engine{gin.New(), admin.NewAdminHttpAdminHandler(path), wx.NewWxHttpAdminHandler(path)}
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

	weiXin := r.Group("/wx")
	{
		//获取所有biz和name信息
		weiXin.GET("/biz", r.FindBiz)
		//获取点赞等接口
		weiXin.GET("/api", r.FindApi)
		//提交点赞和阅读量数据接口
		weiXin.POST("/post", r.PostData)
	}
	return r.Engine
}
