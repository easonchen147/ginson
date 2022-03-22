package middleware

import (
	"ginson/pkg/constant"
	"ginson/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constant.TraceIdKey, strings.ReplaceAll(utils.GetUuidV4(), "-", ""))
	}
}
