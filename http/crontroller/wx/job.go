package wx

import (
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
		Id        string `json:"id"`         //主键id  ArticleId
		Title     string `json:"title"`      //标题
		Text      string `json:"text"`       //正文
		TextStyle string `json:"text_style"` //带样式的正文
		Biz       string `json:"biz"`        //biz
		//Ctime     time.Time `json:"ctime"`      //入库时间
		//Mtime     time.Time `json:"mtime"`      //修改时间
		Ptime    time.Time `json:"ptime"`    //发布时间
		Author   string    `json:"author"`   //作者
		Original int       `json:"original"` //原创
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

}
