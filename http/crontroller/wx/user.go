package wx

import (
	"comadmin/model/wx"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"net/http"
)

func (h HttpWxHandler) OwnWx(c app.GContext) {
	g := app.G{c}
	xin := wx.WeiXin{}
	m := make(map[string]interface{})

	if value, exists := g.Get("anonymous"); exists && value.(bool) {
		queryMap := map[string]interface{}{" `default` = ": 1}
		biz, count := h.logic.FindOne(xin, queryMap, 20, 0, CtimeDesc)
		m["list"] = biz
		m["count"] = count
	}
	g.Json(http.StatusOK, e.Success, m)
	return
}
