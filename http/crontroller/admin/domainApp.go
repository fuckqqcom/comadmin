package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"log"
	"net/http"
)

/**
创建应用app
*/
func (h HttpHandler) CreateDomainApp(c app.GContext) {
	g := app.G{c}
	type P struct {
		Name string `json:"name"  binding:"required"`
		Did  string `json:"did" binding:"required"`
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	domain := &admin.Domain{Id: p.Did}
	code = h.logic.FindById(domain)
	if code != e.Success {
		g.Json(http.StatusOK, code, "")
		return
	}
	domainApp := admin.DomainApp{Name: p.Name, Did: p.Did, Id: utils.EncodeMd5(p.Name + p.Did), Status: 1}
	code = h.logic.Create(domainApp)
	g.Json(http.StatusOK, code, "")
	return
}

/**
更新应用app(主要是更新名字)
*/

func (h HttpHandler) UpdateDomainApp(c app.GContext) {
	g := app.G{c}
	type P struct {
		Name   string `json:"name"`
		Id     string `json:"id" binding:"required"`
		Status int    `json:"status"`
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	cols := make([]string, 0)
	domainApp := admin.DomainApp{Id: p.Id}
	if p.Name != "" {
		domainApp.Name = p.Name
		cols = append(cols, "name")

	}
	if p.Status != 0 {
		domainApp.Status = p.Status
		cols = append(cols, "status")

	}
	code = h.logic.Update(domainApp, cols)
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpHandler) DeleteDomainApp(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id string `json:"id"  binding:"required"`
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "deleteDomainApp") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return

	}
	domain := admin.DomainApp{Id: p.Id}
	code = h.logic.Delete(domain)
	g.Json(http.StatusOK, code, "")
	return
}

/**
  查找，只能查找自己当前did下面的app
	比如 当前用户操作(属于这个域的管理员),选择列表的时候出现app列表
*/
func (h HttpHandler) FindDomainApp(c app.GContext) {
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
	domain := admin.DomainApp{Id: p.Id, Name: p.Name, Status: p.Status, Did: did.(string)}

	list, count := h.logic.Find(domain, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}