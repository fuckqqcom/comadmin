package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
)

/**
我们添加wx
*/
func (h HttpWxHandler) AddWx(c app.GContext) {
	type params struct {
		Biz  string `json:"biz"`  //公号biz
		Name string `json:"name"` //公号name
		Desc string `json:"desc"`
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	xin := wx.WeiXin{}

	if p.Name != "" {
		xin.Name = p.Name
	}
	if p.Biz != "" {
		xin.Biz = p.Biz
	}
	if p.Desc != "" {
		xin.Biz = p.Biz
	}
	xin.Forbid = 1
	code = h.logic.Create(xin)
	g.Json(http.StatusOK, code, "")
	return

}

//禁用公号
func (h HttpWxHandler) ForBidWx(c app.GContext) {
	type params struct {
		Biz string `json:"biz"  binding:"required" ` //公号biz
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "createDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	xin := wx.WeiXin{Biz: p.Biz, Forbid: -1}
	cols := []string{"forbid"}

	code = h.logic.Update(xin, cols)
	g.Json(http.StatusOK, code, "")
	return
}

//用户提交公号
func (h HttpWxHandler) UserAddWx(c app.GContext) {
	type params struct {
		Uid  string `json:"uid"  binding:"required"`   //用户id
		Name string `json:"name"  binding:"required" ` //公号name
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "userAddWx") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	xin := wx.UserWx{Uid: p.Uid, Name: p.Name}
	xin.Status = 0
	code = h.logic.Create(xin)
	g.Json(http.StatusOK, code, "")
	return
}

//获取所有biz信息
func (h HttpWxHandler) FindBiz(c app.GContext) {

	type params struct {
		MobileId string `json:"mobile_id"  binding:"required"` //手机id
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "userAddWx") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	biz, count := h.logic.FindBiz(p.MobileId)
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
	type params struct {
		Biz        string `json:"biz"`
		ArticleId  string `json:"article_id"`
		ReadCount  int    `json:"read_count"`
		ThumbCount int    `json:"thumb_count"`
	}
	g := app.G{c}

	var p params
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

//前端查询接口
func (h HttpWxHandler) FindDetail(c app.GContext) {
	g := app.G{c}

	var p wx.WeiXinParams
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	if p.Pn <= 1 {
		p.Pn = 1
	}

	if p.Ps <= 0 || p.Ps > 50 {
		p.Ps = 50
	}

	list, count := h.logic.Find(p, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}

/**
入队列
*/
func (h HttpWxHandler) AddQueue(c app.GContext) {
	g := app.G{c}

	var p wx.WeiXinList
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	code = h.logic.Create(p)
	g.Json(http.StatusOK, code, "")
	return
}

/**
获取队列数据
*/

func (h HttpWxHandler) PopQueue(c app.GContext) {

	type params struct {
		Num int `json:"num"`
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	xinList := wx.WeiXinList{}
	list, count := h.logic.Find(xinList, p.Num, 0)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}

/**
获取近七天发布的文章
*/

func (h HttpWxHandler) Nearly7Day(c app.GContext) {
	type params struct {
		Pn  int    `json:"pn"`
		Biz string `json:"biz"`
		Ps  int    `json:"ps"`
	}

	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	if p.Pn <= 1 {
		p.Pn = 1
	}

	if p.Ps <= 0 || p.Ps > 50 {
		p.Ps = 200
	}
	w := wx.WeiXinList{Biz: p.Biz}
	list, count := h.logic.FindList(w, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
