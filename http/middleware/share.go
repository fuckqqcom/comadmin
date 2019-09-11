package middleware

import (
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Share() gin.HandlerFunc {
	return func(c app.GContext) {
		g := app.G{c}
		ip := g.ClientIP()
		ipList := [...]string{"127.0.0.1"}
		f := false
		for _, v := range ipList {
			if ip == v {
				f = true
				break
			}
		}
		if !f {
			g.Json(http.StatusOK, e.Forbid, "")
			g.Abort()
			return
		}
		g.Next()
	}
}
