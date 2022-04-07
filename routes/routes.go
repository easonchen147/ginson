package routes

import (
	"ginson/app/tool"
	"ginson/app/user"
	"ginson/app/wxmini"
	"ginson/pkg/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", api.NewHandler().Index)

	r := server.Group("/api")
	{
		user.BindUserRoutes(r.Group("/user"))
		wxmini.BindWxMiniRoutes(r.Group("/wx-mini"))
		tool.BindToolRoutes(r.Group("/tool"))
	}
}
