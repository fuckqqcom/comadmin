package middleware

import (
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	webtoken "comadmin/pkg/jwt"
	"github.com/gin-gonic/gin"

	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c app.GContext) {
		var code int
		var data interface{}

		code = e.Success
		token := c.Request.Header.Get("token")
		if token == "" {
			code = e.Unauthorized
		} else {
			//根据实际需要设置数据
			if claims, code := webtoken.ParseToken(token); code == e.Success {
				c.Set("userName", claims.Username)
				c.Set("userId", claims.Id)
				c.Set("isAdmin", claims.IsAdmin)
				c.Set("isRoot", claims.IsRoot)
			}
		}
		if code != e.Success {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
