package middleware

import (
	"ginson/foundation/util"
	"ginson/pkg/constant"
	"github.com/gin-gonic/gin"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constant.TraceIdKey, util.GetNanoId())
	}
}
