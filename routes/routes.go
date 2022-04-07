package routes

import (
	"ginson/app/service/tool"
	"ginson/app/service/user"
	"ginson/app/service/wxmini"
	"ginson/pkg/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", api.NewBaseHandler().Index)

	r := server.Group("/api")
	{
		user.BindUserRoutes(r.Group("/user"))
		wxmini.BindWxMiniRoutes(r.Group("/wx-mini"))
		tool.BindToolRoutes(r.Group("/tool"))
	}
}
