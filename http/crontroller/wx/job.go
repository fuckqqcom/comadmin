package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"fmt"
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
		Ptime     string `json:"ptime" binding:"required" `       //发布时间
		Author    string `json:"author" binding:"required" `      //作者
		From      string `json:"from"  binding:"required"`
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
	id := utils.EncodeMd5(p.Ip)
	if !h.logic.Exist(&job, nil) {
		job.Id = id
		job.Status = 1 //在线状态
		job.Count = 1
		code = h.logic.Create(job)
	} else {
		code = e.ExistError

	}
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
	job := wx.Job{Id: p.Id, IP: p.Ip}

	count := h.logic.Get(&job, []string{"count"}, nil)
	bean := count.(*wx.Job)
	if bean.Count == 0 {
		code = e.JobError
		g.Json(http.StatusOK, code, "")
		return
	}
	fmt.Println(count)
	//这一层在dao层做
	cols := []string{"count", "etime", "status"}
	//job := wx.Job{Id: p.Id, IP: p.Ip, Etime: time.Now().Local(), Count: count + 1, Status: -1}
	job.Etime = time.Now().Local()
	job.Count = bean.Count + 1
	job.Status = -1
	code = h.logic.Update(job, cols, nil)
	g.Json(http.StatusOK, code, "")
	return
}
