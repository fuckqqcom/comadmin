package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
)

/**
创建domain
*/
func (h HttpHandler) CreateDomain(c app.GContext) {
	g := app.G{c}
	type P struct {
		Name string `json:"name"  binding:"required"`
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return

	}
	domain := admin.Domain{Name: p.Name, Id: utils.EncodeMd5(p.Name)}
	code = h.logic.Create(domain)
	g.Json(http.StatusOK, code, "")
	return
}

/**
通过id去删除domain
*/
func (h HttpHandler) DeleteDoDomain(c app.GContext) {
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
	domain := admin.Domain{Id: p.Id}
	code = h.logic.Delete(domain)
	g.Json(http.StatusOK, code, "")
	return
}

/**
通过id更新数据
*/

func (h HttpHandler) UpdateDomain(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id" binding:"required"`
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.Domain{Id: p.Id}

	if p.Name != "" {
		domain.Name = p.Name
	}
	if p.Status != 0 {
		domain.Status = p.Status
	}
	code = h.logic.Update(domain)
	g.Json(http.StatusOK, code, "")
	return
}

/**
不管是通过id还是name查询
id精确查询 name模糊查询,返回的都是一个数组
*/
func (h HttpHandler) FindDomainArgs(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id" binding:"required"`
		Name   string `json:"name"`
		Status int    `json:"status"`
		Pn     int    `json:"pn"`
		Ps     int    `json:"ps"`
	}

	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.Domain{Id: p.Id, Name: p.Name, Status: p.Status}

	if domain.Name == "" && domain.Status == 0 {
		code = e.ParamLose
		g.Json(http.StatusOK, code, "")
		return
	}
	list, count := h.logic.Find(domain, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
