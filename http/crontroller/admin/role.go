package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"log"
	"net/http"
)

func (h HttpAdminHandler) CreateRole(c app.GContext) {
	g := app.G{c}
	type P struct {
		Name string `json:"name"  binding:"required"`
		Did  string `json:"did" binding:"required"` //给某个域
		Aid  string `json:"aid" binding:"required"` //某个app创建角色
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	//查询这个域下的app是否存在
	domain := admin.DomainApp{Id: p.Aid, Did: p.Did}

	code = h.logic.FindById(domain)
	if code != e.Success {
		g.Json(http.StatusOK, code, "")
		return
	}
	role := admin.Role{Name: p.Name, Id: utils.EncodeMd5(utils.StringJoin(p.Did, p.Aid, p.Name))}
	appRole := admin.DomainAppRole{Id: utils.EncodeMd5(utils.StringJoin(p.Did, p.Aid, p.Name)), Did: p.Did, Aid: p.Aid, Rid: role.Id}

	//TODO  可以改成channel  goroutine
	code = h.logic.Create(appRole)
	code = h.logic.Create(role)
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpAdminHandler) DeleteRole(c app.GContext) {
	g := app.G{c}

	type P struct {
		Id string `json:"id"  binding:"required"`
	}

	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "deleteDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return

	}
	//删除主表中的id
	domain := admin.Role{Id: p.Id}
	//删除关系表中的数据
	role := admin.DomainAppRole{Rid: p.Id}
	code = h.logic.Delete(role)
	code = h.logic.Delete(domain)
	g.Json(http.StatusOK, code, "")
	return
}

//只更新角色名称
func (h HttpAdminHandler) UpdateRole(c app.GContext) {
	g := app.G{c}
	type P struct {
		Name string `json:"name"  binding:"required"`
		Id   string `json:"id"     binding:"required"`
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") || p.Name == "" {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	cols := []string{"name"}
	role := admin.Role{Id: p.Id}

	code = h.logic.Update(role, cols)
	g.Json(http.StatusOK, code, "")
	return
}

//禁用角色
func (h HttpAdminHandler) ForbidRole(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id" binding:"required"`
		Status int    `json:"status"  binding:"required" `
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") || p.Status == 0 {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	cols := []string{"status"}
	role := admin.Role{Id: p.Id}
	appRole := admin.DomainAppRole{Rid: p.Id, Status: p.Status}

	//同时两个表都更新
	code = h.logic.Update(appRole, cols)
	code = h.logic.Update(role, cols)
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpAdminHandler) FindRole(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Status int    `json:"status"`
		Pn     int    `json:"pn"`
		Ps     int    `json:"ps"`
	}

	did, exists := c.Get("did")
	did = "71abfd41229b11e2f431750af5f7693f"
	if !exists {
		log.Printf("获取did error %s", did)
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	role := admin.Role{Id: p.Id, Name: p.Name, Status: p.Status}

	list, count := h.logic.Find(role, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
