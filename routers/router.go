package routers

import (
	"ginson/pkg/resp"
	"ginson/routers/api/tool"
	"ginson/routers/api/user"
	"ginson/routers/api/wxmini"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", resp.NewHandler().Index)

	r := server.Group("/api")
	{
		user.RegisterUserRouters(r.Group("/user"))
		wxmini.RegisterWxMiniRouters(r.Group("/wx-mini"))
		tool.RegisterToolRouters(r.Group("/tool"))
	}
}
