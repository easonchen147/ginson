package tool

import (
	"fmt"
	"ginson/pkg/conf"
	"ginson/pkg/resp"
	"ginson/service/tool"
	"github.com/easonchen147/foundation/log"
	"github.com/easonchen147/foundation/util"
	"image/color"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type handler struct {
	*resp.Handler
	service *tool.Service
}

func newHandler() *handler {
	return &handler{
		Handler: resp.NewHandler(),
		service: tool.NewService(),
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

func (c *handler) imageFaceLocation(ctx *gin.Context) {
	filePath, err := c.saveUploadImageFile(ctx)
	if err != nil {
		c.FailedWithErr(ctx, err)
		return
	}

	//TODO scan image for sensitive

	var result map[string]interface{}
	result, err = c.service.FaceLocation(ctx, filePath)
	defer os.Remove(filePath) // 图片识别完毕后，删除对应文件
	if err != nil {
		c.FailedWithErr(ctx, err)
		return
	}

	c.SuccessData(ctx, result)
}

func (c *handler) saveUploadImageFile(ctx *gin.Context) (string, error) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		return "", err
	}

	filePath := conf.ExtConf().UploadImagePath + string(os.PathSeparator) + util.GetNanoId() + string(os.PathSeparator) + fmt.Sprintf("%v", time.Now().Unix()) + filepath.Ext(header.Filename)
	log.Info(ctx, "image face location, by file path: %v", filePath)

	out, err := os.Create(filePath)
	if err != nil {
		log.Error(ctx, "create file failed, error: %v", err)
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Error(ctx, "copy file failed, error: %v", err)
		return "", err
	}
	return filePath, err
}
