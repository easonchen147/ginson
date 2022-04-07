package tool

import (
	"github.com/gin-gonic/gin"
)

func BindToolRoutes(group *gin.RouterGroup) {
	handler := NewHandler()
	group.GET("/get-qr-code", handler.GetQrCode)
	group.GET("/get-screenshot", handler.GetScreenShot)
}
