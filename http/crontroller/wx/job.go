package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
	"time"
)

/**
接受远端传输数据(目前只支持http通信)
*/

func (h HttpWxHandler) AddDetail(c app.GContext) {
	type params struct {
		Id        string `json:"id"  binding:"required" `         //主键id  ArticleId
		Title     string `json:"title"  binding:"required" `      //标题
		Text      string `json:"text"  binding:"required" `       //正文
		TextStyle string `json:"text_style"  binding:"required" ` //带样式的正文
		Biz       string `json:"biz"  binding:"required" `        //biz
		//Ctime     time.Time `json:"ctime"`      //入库时间
		//Mtime     time.Time `json:"mtime"`      //修改时间
		Ptime  string `json:"ptime" binding:"required" `  //发布时间
		Author string `json:"author" binding:"required" ` //作者
		From   string `json:"from"  binding:"required"`
		//Original int       `json:"original" binding:"required"` //原创
		//WordCloud string    `json:"word_cloud"` //词云数据
		//Summary   string    `json:"summary"`    //摘要
		//Forbid    int       `json:"forbid"`     //公号是否被微信官方搞事了
	}

	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	const t = "2006-01-02 15:04:05"

	detail := wx.WeiXinDetail{Id: p.Id, Title: p.Title, Text: p.Text, TextStyle: p.TextStyle, Biz: p.Biz,
		Ptime: utils.Str2Time(t, p.Ptime), Author: p.Author, Forbid: 1, Ctime: time.Now().Local(), Mtime: time.Now().Local()}
	h.logic.Create(detail)
	g.Json(http.StatusOK, code, "")
	return

}

/**
注册job任务
*/
func (h HttpWxHandler) RegisterJob(c app.GContext) {
	type params struct {
		Id string `json:"id" ` //第一次注册的时候 id为kong
		Ip string `json:"ip" binding:"required" `
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	job := wx.Job{Id: p.Id, IP: p.Ip}
	id, code := h.logic.Register(job)
	g.Json(http.StatusOK, code, id)
	return

}

func (h HttpWxHandler) UpdateJob(c app.GContext) {
	type params struct {
		Id string `json:"id" binding:"required" `
		Ip string `json:"ip" binding:"required" `
	}
	g := app.G{c}

	var p params
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	count := h.logic.FindCountByIdAndIp(p.Id, p.Ip)
	if count == 0 {
		code = e.JobError
		g.Json(http.StatusOK, code, "")
		return
	}
	cols := []string{"count", "etime", "status"}
	job := wx.Job{Id: p.Id, IP: p.Ip, Etime: time.Now().Local(), Count: count + 1, Status: -1}
	code = h.logic.Update(job, cols)
	g.Json(http.StatusOK, code, "")
	return
}
