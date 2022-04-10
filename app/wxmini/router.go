package wxmini

import (
	"github.com/gin-gonic/gin"
)

func BindWxMiniRoutes(group *gin.RouterGroup) {
	wxMini := newHandler()
	group.POST("/login", wxMini.WxMiniLogin)
}
