package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
)

//ExportPdf
//middleware
func (h HttpWxHandler) DownLoad(c app.GContext) {
	g := app.G{c}
	code := e.Success

	var p wx.BizPdf
	if !utils.CheckError(c.ShouldBindJSON(&p), "GetBiz") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	//查询条件
	queryMap := make(map[string]interface{})
	if p.Biz != "" {
		queryMap["biz"] = p.Biz
	} else {
		queryMap["ids"] = p.Ids
	}
	pdf := utils.Pdf{Path: "./", Title: "1", Content: "hahahaha"}
	pdf.HtmlToPdf()
}
