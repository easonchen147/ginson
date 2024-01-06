package routers

import (
	"ginson/pkg/resp"
	"ginson/routers/handler/tool"
	"ginson/routers/handler/user"
	"ginson/routers/handler/wxmini"
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
