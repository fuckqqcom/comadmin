package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
)

//获取所有biz信息
func (h HttpWxHandler) FindBiz(c app.GContext) {
	g := app.G{c}

	biz, count := h.logic.FindBiz()
	m := make(map[string]interface{})
	m["list"] = biz
	m["count"] = count
	g.Json(http.StatusOK, e.Success, m)
	return
}

//获取点赞和阅读接口
func (h HttpWxHandler) FindApi(c app.GContext) {
	g := app.G{c}

	biz, count := h.logic.FindApi()
	m := make(map[string]interface{})
	m["list"] = biz
	m["count"] = count
	g.Json(http.StatusOK, e.Success, m)
	return
}

func (h HttpWxHandler) PostData(c app.GContext) {
	type param struct {
		Biz        string `json:"biz"`
		ArticleId  string `json:"article_id"`
		ReadCount  int    `json:"read_count"`
		ThumbCount int    `json:"thumb_count"`
	}
	g := app.G{c}

	var p param
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	wxCount := wx.WeiXinCount{Biz: p.Biz, ArticleId: p.ArticleId, ReadCount: p.ReadCount, ThumbCount: p.ThumbCount}
	code = h.logic.PostData(wxCount)
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpWxHandler) FindDetail(c app.GContext) {
	g := app.G{c}

	var p wx.WeiXinParams
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	list, count := h.logic.Find(p, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
