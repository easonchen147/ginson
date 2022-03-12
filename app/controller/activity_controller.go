package controller

import (
	"fmt"
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
	log.Info(fmt.Sprintf("%s come in", ctx.Value("openId")))

	c.Success(ctx, nil)
}
