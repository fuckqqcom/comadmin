package middleware

import (
	"comadmin/pkg/app"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
查询用户是否
*/

func UserExist() gin.HandlerFunc {
	return func(c app.GContext) {

		fmt.Println("userExist......")
		c.Next()
	}
}
