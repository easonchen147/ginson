package util

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"time"

	"ginson/pkg/constant"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func MD5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(str []byte) string {
	h := sha1.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

// FileHash 计算文件hash
func FileHash(reader io.Reader, tp string) string {
	var result []byte
	var h hash.Hash
	if tp == "md5" {
		h = md5.New()
	} else {
		h = sha1.New()
	}
	if _, err := io.Copy(h, reader); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(result))
}

// GetUuid 生成uuid
func GetUuid() string {
	var u uuid.UUID
	var err error
	for i := 0; i < 3; i++ {
		u, err = uuid.NewUUID()
		if err == nil {
			return u.String()
		}
	}
	return ""
}

// GetUuidV4 生成uuid v4
func GetUuidV4() string {
	var u uuid.UUID
	var err error
	for i := 0; i < 3; i++ {
		u, err = uuid.NewRandom()
		if err == nil {
			return u.String()
		}
	}
	return ""
}

// ParseDate 转时间格式 yyyy-MM-dd
func ParseDate(date string) time.Time {
	result, err := time.ParseInLocation(constant.DateFormat, date, time.Local)
	if err != nil {
		return time.Time{}
	}
	return result
}

// ParseDateTime 转时间格式 yyyy-MM-dd HH:mm:ss
func ParseDateTime(dateTime string) time.Time {
	result, err := time.ParseInLocation(constant.DateTimeFormat, dateTime, time.Local)
	if err != nil {
		return time.Time{}
	}
	return result
}

// FormatDate 转时间格式 yyyy-MM-dd
func FormatDate(date time.Time) string {
	return date.Format(constant.DateFormat)
}

// FormatDateTime 转时间格式 yyyy-MM-dd HH:mm:ss
func FormatDateTime(dateTime time.Time) string {
	return dateTime.Format(constant.DateTimeFormat)
}

// CopyGinCtx 复制ginCtx
func CopyGinCtx(ctx *gin.Context) context.Context {
	return ctx.Copy()
}

// CopyCtx 拷贝Context，判断如果是ginCtx则需要拷贝
func CopyCtx(ctx context.Context) context.Context {
	switch ctx.(type) {
	case *gin.Context:
		return ctx.(*gin.Context).Copy()
	default:
		return ctx
	}
}

const nanoIdAlphbet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GetNanoId 获取32位的nanoId
func GetNanoId() string {
	id, err := gonanoid.Generate(nanoIdAlphbet, 32)
	if err != nil {
		return ""
	}
	return id
}
