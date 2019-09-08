package router

import (
	"comadmin/http/crontroller/admin"
	"comadmin/http/crontroller/wx"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func NewRouter(path string) *gin.Engine {

	r := &Engine{gin.New(), admin.NewAdminHttpAdminHandler(path), wx.NewWxHttpAdminHandler(path)}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.JWT())
	g := r.Group("/v1")
	domain := g.Group("/d")
	{
		//域
		domain.POST("/addDomain", r.AddDomain)
		domain.POST("/deleteDomain", r.DeleteDoDomain)
		domain.GET("/domains", r.Domains)
		domain.POST("/editDomain", r.EditDomain)
		////app
		domain.POST("/createApp", r.AddDomainApp)
		domain.POST("/updateApp", r.UpdateDomainApp)
		domain.POST("/findApp", r.FindDomainApp)
		//domain.POST("/deleteApp", r.DeleteDomainApp)
	}

	//用户操作
	//user := g.Group("/user")
	//user.POST("/register", r.Register) //用户注册
	//user.POST("/login", r.Login)       //用户登录
	//{
	//	user.POST("/createAdmin", r.RegisterAdminUser)
	//	user.POST("/delete", r.DeleteUser)
	//	user.POST("/find", r.FindUser)
	//	user.POST("/updatePhone", r.UpdateUserPhone)
	//	user.POST("/updatePwd", r.UpdateUserPwd)
	//	user.POST("/forbidUser", r.UpdateUserStatus)
	//}

	weiXin := g.Group("/wx")
	{
		/**
		提交数据
		*/
		//提交点赞和阅读量数据接口  readAndThumbCount
		weiXin.POST("/readAndThumbCount", r.ReadAndThumbCount)
		//详情页入库
		weiXin.POST("/addDetail", r.AddDetail)
		//用户添加wx号
		weiXin.POST("/addUserWx", r.UserAddWx)
		//后台添加公号
		weiXin.POST("/addAdminWx", r.AddWx)
		//job注册
		weiXin.POST("/onlineJob", r.OnlineJob)

		//job离线
		weiXin.POST("/offlineJob", r.OfflineJob)
		//入队列(列表数据入库)
		weiXin.POST("/addWxList", r.AddWxList)
		/**
		查询数据
		*/
		//获取所有biz和name信息
		weiXin.GET("/getBiz", r.GetBiz)
		//获取点赞等接口
		weiXin.GET("/getApi", r.GetApi)

		//获取近七天的文章列表
		weiXin.GET("/nearly7Day", r.Nearly7Day)

		//获取队列任务(默认是5个job)
		weiXin.GET("/getTasks", r.GetTasks)

		//查询数据接口
		weiXin.POST("/getDetail", r.GetDetail)

		//更新微信key
		weiXin.POST("/updateKey", r.UpdateKey)

		weiXin.POST("/download", r.DownLoad)

	}
	ginpprof.Wrapper(r.Engine)
	return r.Engine
}
