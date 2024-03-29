package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
)

/**
创建应用app
*/
func (h HttpAdminHandler) AddDomainApp(c app.GContext) {
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

	domain := admin.Domain{Id: p.Did}
	//先查domain是否存在
	exist := h.logic.Exist(&domain, nil)
	if !exist {
		code = e.NotExistError
		g.Json(http.StatusOK, code, "")
		return
	}
	domainApp := admin.DomainApp{Id: utils.EncodeMd5(p.Name + p.Did)}
	if !h.logic.Exist(&domainApp, nil) {
		domainApp.Name = p.Name
		domainApp.Did = p.Did
		domainApp.Status = 1
		code = h.logic.Add(domainApp)
	} else {
		code = e.ExistError
	}
	//domainApp := admin.DomainApp{Name: p.Name, Did: p.Did, Id: utils.EncodeMd5(p.Name + p.Did), Status: 1}
	//code = h.logic.Add(domainApp)
	g.Json(http.StatusOK, code, "")
	return
}

/**
更新应用app(主要是更新名字)
*/

func (h HttpAdminHandler) UpdateDomainApp(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id" binding:"required"`
		Name   string `json:"name"`
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
	queryMap := make(map[string]interface{}) //条件
	queryMap[" id = "] = p.Id
	domainApp := admin.DomainApp{Id: p.Id}
	if p.Name != "" {
		domainApp.Name = p.Name
		cols = append(cols, "name")

	}
	if p.Status != 0 {
		domainApp.Status = p.Status
		cols = append(cols, "status")

	}
	code = h.logic.Update(domainApp, cols, queryMap)
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpAdminHandler) DeleteDomainApp(c app.GContext) {
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
	code = h.logic.Delete(domain, nil)
	g.Json(http.StatusOK, code, "")
	return
}

/**
  角色控制(过滤器中做验证)
  查找，只能查找自己当前did下面的app
	比如 当前用户操作(属于这个域的管理员),选择列表的时候出现app列表
*/
func (h HttpAdminHandler) FindDomainApp(c app.GContext) {
	g := app.G{c}
	type P struct {
		Did    string `json:"did"`
		Id     string `json:"id"`
		Name   string `json:"name"`
		Status int    `json:"status"`
		Pn     int    `json:"pn"`
		Ps     int    `json:"ps"`
	}

	//did, exists := c.Get("did")
	//did = "71abfd41229b11e2f431750af5f7693f"
	//if !exists {
	//	log.Printf("获取did error %s", did)
	//}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.DomainApp{Id: p.Id, Name: p.Name, Status: p.Status, Did: p.Did}

	list, count := h.logic.FindOne(domain, nil, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
