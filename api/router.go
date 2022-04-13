package api

import (
	"ginson/api/tool"
	"ginson/api/user"
	"ginson/api/wxmini"
	"ginson/pkg/resp"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", resp.NewHandler().Index)

	r := server.Group("/api")
	{
		user.BindUserRoutes(r.Group("/user"))
		wxmini.BindWxMiniRoutes(r.Group("/wx-mini"))
		tool.BindToolRoutes(r.Group("/tool"))
	}
}
