package middleware

import (
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
匿名用户只能看前几页数据
*/

func PageForBid() gin.HandlerFunc {
	return func(c app.GContext) {
		g := app.G{c}
		if _, ok := c.Get("userId"); ok {
			g.Next()
		} else {
			g.Json(http.StatusOK, e.NoLogin, "")
			g.Abort()
		}

	}
}
