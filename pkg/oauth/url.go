package oauth

import (
	"fmt"
	"net/url"
	"strings"
)

type UrlHelper struct {
	baseUrl string
	params  url.Values
}

func NewUrlHelper(baseUrl string) *UrlHelper {
	u, err := url.ParseRequestURI(baseUrl)
	helper := &UrlHelper{}
	if err != nil {
		return helper
	}

	urls := strings.SplitN(u.String(), "?", 2)
	helper.baseUrl = urls[0]
	helper.params = u.Query()
	return helper
}

func (u *UrlHelper) AddParam(key string, value interface{}) *UrlHelper {
	if key == "" {
		return u
	}
	u.params.Add(key, fmt.Sprint(value))
	return u
}

func (u *UrlHelper) Build() string {
	if u.baseUrl == "" {
		return ""
	}
	if len(u.params) == 0 {
		return u.baseUrl
	}
	return u.baseUrl + "?" + u.params.Encode()
}
