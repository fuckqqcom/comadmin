package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
)

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

	domain := admin.Domain{Id: p.Did}
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
