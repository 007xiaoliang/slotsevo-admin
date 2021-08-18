package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"slotsevo-admin/global"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				global.SlotsLog.DPanic(fmt.Sprintf("%s", r))
			}
		}()
		c.Next()
	}
}
