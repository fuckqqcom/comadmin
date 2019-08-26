package main

import (
	"comadmin/model/admin"
	"comadmin/model/wx"
	"comadmin/pkg/config"
)

func main() {
	config.NewConfig("config/config.json")

	config.EngDb.Sync2(new(admin.Domain))
	config.EngDb.Sync2(new(admin.DomainApp))
	config.EngDb.Sync2(new(admin.User))
	config.EngDb.Sync2(new(admin.DomainAppUser))

	config.EngDb.Sync2(new(wx.WeiXin))
	config.EngDb.Sync2(new(wx.Api))
	config.EngDb.Sync2(new(wx.WeiXinCount))
	config.EngDb.Sync2(new(wx.WeiXinList))
	config.EngDb.Sync2(new(wx.UserWx))

	//角色相关
	//config.EngDb.Sync2(new(admin.Role))
	//config.EngDb.Sync2(new(admin.DomainAppRole))

}
