package crontroller

import (
	"comadmin/model/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h HttpHandler) CreateDomain(c *gin.Context) {
	name := c.Query("name")
	domain := admin.Domain{Name: name}
	code := h.logic.Create(domain)
	c.JSON(http.StatusOK, code)
	return
}

func (h HttpHandler) DeleteDoDomain(c *gin.Context) {
	id := c.Query("id")
	domain := admin.Domain{Id: id}
	code := h.logic.Delete(domain)
	c.JSON(http.StatusOK, code)
	return
}
