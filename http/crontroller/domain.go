package crontroller

import (
	"comadmin/model/admin"
	"comadmin/pkg/app"
	"comadmin/tools/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h HttpHandler) CreateDomain(c app.GContext) {
	g := app.G{c}
	name := c.Query("name")
	domain := admin.Domain{Name: name, Id: utils.EncodeMd5(name)}
	code := h.logic.Create(domain)
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpHandler) DeleteDoDomain(c *gin.Context) {
	g := app.G{c}
	id := c.Query("id")
	domain := admin.Domain{Id: id}
	code := h.logic.Delete(domain)
	g.Json(http.StatusOK, code, "")
	return
}
