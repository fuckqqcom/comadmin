package main

import (
	"comadmin/model/admin"
	"comadmin/pkg/config"
)

func main() {
	config.NewConfig("config/config.json")

	config.EngDb.Sync2(new(admin.Domain))
	//config.EngDb.Sync2(new(admin.DomainApp))

	//角色相关
	//config.EngDb.Sync2(new(admin.Role))
	//config.EngDb.Sync2(new(admin.DomainAppRole))

}
