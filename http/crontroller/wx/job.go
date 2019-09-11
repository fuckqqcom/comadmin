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
		Id        string    `json:"id"  binding:"required" `         //主键id  ArticleId
		Title     string    `json:"title"  binding:"required" `      //标题
		Text      string    `json:"text"  binding:"required" `       //正文
		TextStyle string    `json:"text_style"  binding:"required" ` //带样式的正文
		Biz       string    `json:"biz"  binding:"required" `        //biz
		Ptime     time.Time `json:"ptime" binding:"required" `       //发布时间
		Author    string    `json:"author" binding:"required" `      //作者
		From      string    `json:"from"  binding:"required"`
		Url       string    `json:"url"  binding:"required"` //url连接
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
		Ptime: p.Ptime, Author: p.Author, Forbid: 1, Ctime: time.Now().Local(), Mtime: time.Now().Local(), Url: p.Url, From: p.From}
	h.logic.Add(detail)
	g.Json(http.StatusOK, code, "")
	return

}

/**
注册job任务
*/
func (h HttpWxHandler) OnlineJob(c app.GContext) {
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
	job := wx.Job{Id: p.Id, Ip: p.Ip}
	id := utils.EncodeMd5(p.Ip)

	/**
	先查询是否存在,不存在就创建
	存在就更新字段
	*/
	m := make(map[string]interface{})

	inter := h.logic.Get(&job, []string{" `count` ", "job_count", " `interval` "}, nil)
	if inter == nil {
		job.Id = id
		job.Status = 1 //在线状态
		job.Count = 1
		m["job_count"] = 3
		m["interval"] = 10
		code = h.logic.Add(job)
	} else {
		bean := inter.(*wx.Job)
		cols := []string{"count", "etime", "status"}
		job.Etime = time.Now().Local()
		job.Count = bean.Count + 1
		job.Status = 1
		m["job_count"] = bean.JobCount
		m["interval"] = bean.Interval

		queryMap := make(map[string]interface{})
		queryMap[" id = "] = id
		code = h.logic.Update(job, cols, queryMap)
	}
	m["id"] = id
	g.Json(http.StatusOK, code, m)
	return

}

func (h HttpWxHandler) OfflineJob(c app.GContext) {
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
	job := wx.Job{Id: p.Id, Ip: p.Ip}
	//这一层在dao层做
	cols := []string{"etime", "status"}
	job.Etime = time.Now().Local()
	job.Status = -1
	queryMap := make(map[string]interface{})
	queryMap[" id = "] = p.Id
	code = h.logic.Update(job, cols, queryMap)
	g.Json(http.StatusOK, code, "")
	return
}
