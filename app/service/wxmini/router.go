package wxmini

import (
	"github.com/gin-gonic/gin"
)

func BindWxMiniRoutes(group *gin.RouterGroup) {
	wxMini := NewHandler()
	group.POST("/login", wxMini.WxMiniLogin)
}
