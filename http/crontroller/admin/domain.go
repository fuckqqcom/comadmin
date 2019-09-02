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
func (h HttpAdminHandler) AddDomain(c app.GContext) {
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
	domain := admin.Domain{Name: p.Name, Id: utils.EncodeMd5(p.Name), Status: 1}
	code = h.logic.Add(domain)
	g.Json(http.StatusOK, code, "")
	return
}

/**
通过id去删除domain
*/
func (h HttpAdminHandler) DeleteDoDomain(c app.GContext) {
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
	domain := admin.Domain{}
	queryMap := map[string]interface{}{" id = ": p.Id}
	code = h.logic.Delete(domain, queryMap)
	g.Json(http.StatusOK, code, "")
	return
}

/**
通过id更新数据
*/

func (h HttpAdminHandler) EditDomain(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id" binding:"required"`
		Name   string `json:"name"`
		Status int    `json:"status"`
	}

	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "editDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.Domain{}
	cols := make([]string, 0)
	if p.Name != "" {
		domain.Name = p.Name
		cols = append(cols, "name")
	}
	if p.Status != 0 {
		domain.Status = p.Status
		cols = append(cols, "status")
	}

	if len(cols) == 0 {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	queryMap := map[string]interface{}{" id = ": p.Id}
	code = h.logic.Update(domain, cols, queryMap)
	g.Json(http.StatusOK, code, "")
	return
}

/**
不管是通过id还是name查询
id精确查询 name模糊查询,返回的都是一个数组
*/
func (h HttpAdminHandler) Domains(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id"`
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
	domain := admin.Domain{}

	queryMap := make(map[string]interface{}, 0)

	if p.Id != "" {
		queryMap["id = "] = p.Id
	}

	if p.Name != "" {
		// NAME LIKE '%新%'
		queryMap[" name like "] = "%" + p.Name + "%"
	}
	if p.Status != 0 {
		queryMap["status = "] = p.Status
	}

	list, count := h.logic.FindOne(domain, queryMap, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
