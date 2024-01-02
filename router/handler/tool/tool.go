package tool

import (
	"ginson/pkg/resp"
	"github.com/easonchen147/foundation/log"
	"image/color"
	"net/url"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type handler struct {
	*resp.Handler
}

func newHandler() *handler {
	return &handler{
		Handler: resp.NewHandler(),
	}
}

func RegisterToolRouters(group *gin.RouterGroup) {
	handler := newHandler()
	group.GET("/get-qr-code", handler.GetQrCode)
	group.GET("/get-screenshot", handler.GetScreenShot)
}

func (c *handler) GetQrCode(ctx *gin.Context) {
	data := ctx.Query("data")

	q, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		c.FailedWithErr(ctx, err)
		return
	}

	q.DisableBorder = true
	q.ForegroundColor = color.RGBA{R: 0x33, G: 0x33, B: 0x66, A: 0xff}
	q.BackgroundColor = color.RGBA{R: 0xef, G: 0xef, B: 0xef, A: 0xff}

	err = q.Write(200, ctx.Writer)
	if err != nil {
		c.FailedWithErr(ctx, err)
		return
	}
}

func (c *handler) GetScreenShot(ctx *gin.Context) {
	query := ctx.Query("url")

	chromedp.WithLogf(func(s string, i ...interface{}) {
		log.Info(ctx, s, i)
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

func (c *handler) fullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Emulate(device.IPhoneXR),
		chromedp.Navigate(url),
		//chromedp.EvaluateAsDevTools(`p = document.querySelector("#hotsearch-refresh-btn > span");p.innerText="我调整了标题";`, nil),
		chromedp.FullScreenshot(res, quality),
	}
}
