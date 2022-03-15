package controller

import (
	"ginson/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type ActivityController struct {
	*Controller
}

var activityController = &ActivityController{
	Controller: BaseController,
}

func GetActivityController() *ActivityController {
	return activityController
}

func (c *ActivityController) GetPrize(ctx *gin.Context) {
	//DO SOMETHING
	log.Info("%v come in", ctx.Value("openId"))

	c.Success(ctx, nil)
}

func (c *ActivityController) GetQrCode(ctx *gin.Context) {
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		c.FailedWithErr(ctx, err)
		return
	}
	_, _ = ctx.Writer.Write(png)
}
