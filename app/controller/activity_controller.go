package controller

import (
	"ginson/pkg/log"
	"github.com/gin-gonic/gin"
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
