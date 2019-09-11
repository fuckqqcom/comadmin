package middleware

import (
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	webToken "comadmin/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func Anonymous() gin.HandlerFunc {
	return func(c app.GContext) {
		g := app.G{c}
		token := c.Request.Header.Get("token")

		if token == "" {
			g.Set("anonymous", true)
		} else {
			//根据实际需要设置数据
			if claims, code := webToken.ParseToken(token); code == e.Success && claims != nil {
				g.Set("userName", claims.Username)
				g.Set("userId", claims.Id)
				g.Set("isAdmin", claims.IsAdmin)
				g.Set("isRoot", claims.IsRoot)
				g.Set("anonymous", false)

			} else {
				g.Set("anonymous", true)
			}
		}
		g.Next()
	}
}
