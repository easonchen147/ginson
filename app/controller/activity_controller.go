package controller

import (
	"ginson/pkg/log"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/url"
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

func (c *ActivityController) GetScreenShot(ctx *gin.Context) {
	query := ctx.Query("url")



	chromedp.WithLogf(func(s string, i ...interface{}) {
		log.Info(s, i)
	})

	screenShotUrl := `https://www.baidu.com/`
	if query != "" {
		if _, err := url.Parse(query); err == nil {
			screenShotUrl = query
		}
	}

	cCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(cCtx, c.fullScreenshot(screenShotUrl, 100, &buf)); err != nil {
		c.FailedWithErr(ctx, err)
		return
	}
	_, _ = ctx.Writer.Write(buf)
}

func (c *ActivityController) fullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Emulate(device.IPhoneXR),
		chromedp.Navigate(url),
		//chromedp.EvaluateAsDevTools(`p = document.querySelector("#hotsearch-refresh-btn > span");p.innerText="我是傻逼";`, nil),
		chromedp.FullScreenshot(res, quality),
	}
}
