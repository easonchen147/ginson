package utils

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"ginson/pkg/constant"
	"hash"
	"io"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
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
