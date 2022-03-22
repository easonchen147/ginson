package middleware

import (
	"ginson/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()

		latencyTime := endTime.Sub(startTime) // 执行时间
		reqMethod := c.Request.Method         // 请求方式
		reqPath := c.Request.URL.Path         // 请求路径
		reqQuery := c.Request.URL.RawQuery    // 路径后的参数
		statusCode := c.Writer.Status()       // 状态码

		log.Access(c, "RequestLog",
			zap.Int("code", statusCode),
			zap.String("method", reqMethod),
			zap.String("path", reqPath),
			zap.String("query", reqQuery),
			zap.Duration("cost", latencyTime))
	}
}
