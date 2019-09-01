package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
	"strings"
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
	if p.Biz != "" {
		xin.Biz = p.Biz
	}
	if p.Name != "" {
		xin.Name = p.Name
	}

	if !h.logic.Exist(&xin, nil) {
		xin.Biz = p.Biz
		xin.Forbid = 1
		code = h.logic.Create(xin)
	} else {
		code = e.ExistError
	}
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

	code = h.logic.Update(xin, cols, nil)
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
	xin := wx.UserWx{Name: p.Name}
	//Uid: p.Uid,
	if !h.logic.Exist(&xin, nil) {
		xin.Uid = p.Uid
		xin.Status = 0
		code = h.logic.Create(xin)
	} else {
		code = e.ExistError
	}

	g.Json(http.StatusOK, code, "")
	return
}

//获取所有biz信息
func (h HttpWxHandler) GetBiz(c app.GContext) {

	type params struct {
		MobileId string `json:"mobile_id"  binding:"required"` //手机id
		Ps       int    `json:"ps"`
		Pn       int    `json:"pn"`
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "userAddWx") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	ps, pn := utils.Pagination(p.Ps, p.Pn, 200)

	biz, count := h.logic.List(p.MobileId, ps, pn)
	m := make(map[string]interface{})
	m["list"] = biz
	m["count"] = count
	g.Json(http.StatusOK, e.Success, m)
	return
}

//获取点赞和阅读接口
func (h HttpWxHandler) GetApi(c app.GContext) {
	g := app.G{c}

	api := wx.Api{}
	biz, count := h.logic.List(api, 20, 0)
	m := make(map[string]interface{})
	m["list"] = biz
	m["count"] = count
	g.Json(http.StatusOK, e.Success, m)
	return
}

func (h HttpWxHandler) ReadAndThumbCount(c app.GContext) {
	type params struct {
		Biz        string `json:"biz"  binding:"required"`
		ArticleId  string `json:"article_id"  binding:"required"`
		ReadCount  int    `json:"read_count"`
		ThumbCount int    `json:"thumb_count"`
	}

	//这里优化下，这里组装字段
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "ReadAndThumbCount") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	cols := make([]string, 0)
	if p.ReadCount != 0 {
		cols = append(cols, "read_count")
	}
	if p.ThumbCount != 0 {
		cols = append(cols, "thumb_count")
	}
	colsValue := map[string]interface{}{
		"biz":        p.Biz,
		"article_id": p.ArticleId,
	}

	if len(cols) == 0 {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	wxCount := wx.WeiXinCount{Biz: p.Biz, ArticleId: p.ArticleId}
	query := wxCount
	wxCount.ReadCount = p.ReadCount
	wxCount.ThumbCount = p.ThumbCount
	if h.logic.Exist(&query, nil) {
		code = h.logic.Update(wxCount, cols, colsValue)
	} else {
		code = h.logic.Create(wxCount)
	}

	//wxCount := wx.WeiXinCount{Biz: p.Biz, ArticleId: p.ArticleId, ReadCount: p.ReadCount, ThumbCount: p.ThumbCount}
	//code = h.logic.CreateOrUpdate(wxCount, cols, colsValue)
	g.Json(http.StatusOK, code, "")
	return
}

//前端查询接口
func (h HttpWxHandler) GetDetail(c app.GContext) {
	g := app.G{c}

	var p wx.WeiXinParams
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	ps, pn := utils.Pagination(p.Ps, p.Pn, 200)

	list, count := h.logic.List(p, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}

/**
入队列
*/
func (h HttpWxHandler) AddWxList(c app.GContext) {
	g := app.G{c}

	var p wx.WeiXinList
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "AddWxList") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	//http://mp.weixin.qq.com/s?__biz=MzU3ODE2NTMxNQ==&mid=2247485961&idx=1
	//biz=(\w*).*?mid=(\w*)\w+&idx=(\d+)
	ids := utils.FindBizStr(p.Url)
	if ids != nil {
		p.HashId = utils.EncodeMd5(strings.Join(ids, "_"))
	} else {
		//log  url 提取idx等参数异常
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	m := make(map[string]interface{})
	m["hash_id"] = p.HashId
	//todo 先查询hashId是否存在，存在就不入库，不存在就存在入库
	query := wx.WeiXinList{HashId: p.HashId}
	if !h.logic.Exist(&query, m) {
		code = h.logic.Create(p)
	} else {
		code = e.ExistError
	}

	//code = h.logic.CreateOrDiscard(p, nil, m)
	g.Json(http.StatusOK, code, "")
	return
}

/**
获取队列数据
*/

func (h HttpWxHandler) GetTasks(c app.GContext) {

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
	list, count := h.logic.List(xinList, p.Num, 0)
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

	g := app.G{c}

	var p wx.Nearly7Day
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	/**
	如果biz为空,则查询全部的
	否则查询某一个公号的
	*/
	ps, pn := utils.Pagination(p.Ps, p.Pn, 200)
	w := wx.WeiXinList{}

	if p.Biz != "" {
		w.Biz = p.Biz
	}
	list, count := h.logic.List(w, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
