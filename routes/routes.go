package routes

import (
	"ginson/app/tool"
	"ginson/app/user"
	"ginson/app/wxmini"
	"ginson/pkg/resp"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", resp.NewHandler().Index)

	r := server.Group("/resp")
	{
		user.BindUserRoutes(r.Group("/user"))
		wxmini.BindWxMiniRoutes(r.Group("/wx-mini"))
		tool.BindToolRoutes(r.Group("/tool"))
	}
}
