package user

import (
	"ginson/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func BindUserRoutes(group *gin.RouterGroup) {
	userHandler := newHandler()
	group.Use(middleware.TokenMiddleware())
	group.POST("/get-user-info", userHandler.GetUserInfo)
}
