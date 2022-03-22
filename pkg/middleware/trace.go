package middleware

import (
	"ginson/pkg/constant"
	"ginson/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constant.TraceIdKey, utils.GetUuidV4Simple())
	}
}
