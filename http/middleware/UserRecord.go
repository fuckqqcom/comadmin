package middleware

import (
	"comadmin/pkg/app"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
记录用户最后一次下载记录时间
*/

func DownLoad() gin.HandlerFunc {
	return func(c app.GContext) {

		fmt.Println("download......")
		c.Next()
	}
}
