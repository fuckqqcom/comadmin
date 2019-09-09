package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//ExportPdf
//middleware
func (h HttpWxHandler) DownLoad(c app.GContext) {
	g := app.G{c}

	var p wx.WeiXinParams
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	uid := "xiaohan"
	ps, pn := utils.Pagination(p.Ps, p.Pn, 200)

	//es 倒叙
	list, count := h.logic.FindOne(p, nil, ps, pn, PtimeDesc)
	fmt.Println(list, count)
	type ret struct {
		TextStyle string `json:"text_style"`
		Biz       string `json:"biz"`
		Title     string `json:"title"`
	}
	if p.Type == 1 {
		//导出所有的pdf
		for i := 2; i <= count.(int); i++ {
			fmt.Println(i)
			list, _ := h.logic.FindOne(p, nil, ps, i, PtimeDesc)
			var r ret
			for _, v := range list.([]interface{}) {
				unmarshal := json.Unmarshal([]byte(v.(string)), &r)
				fmt.Println(unmarshal)
				pdf := utils.Pdf{Path: uid + "/" + r.Biz, Title: r.Title, Content: r.TextStyle}
				pdf.HtmlToPdf()
			}
		}
	} else {
		//path 是用户id / biz /xxx
		var r ret
		for _, v := range list.([]interface{}) {
			unmarshal := json.Unmarshal([]byte(v.(string)), &r)
			fmt.Println(unmarshal)
			pdf := utils.Pdf{Path: uid + "/" + r.Biz, Title: r.Title, Content: r.TextStyle}
			pdf.HtmlToPdf()
		}
		//导出当前页pdf
	}

	/**
	导出pdf,做目录压缩
	*/
	zip := utils.Zip(uid, "uid.zip")
	fmt.Println(zip)
}
