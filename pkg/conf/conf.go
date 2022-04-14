package conf

import (
	"fmt"
	"ginson/foundation/cfg"
)

type ExtConfig struct {
	TokenSecret     string `mapstructure:"token_secret"`
	WxMiniAppId     string `mapstructure:"wx_mini_app_id"`
	WxMiniAppSecret string `mapstructure:"wx_mini_app_secret"`
}

var extConfig = &ExtConfig{}

func init() {
	err := cfg.AppConf.LoadExtConfig(&extConfig)
	if err != nil {
		panic(fmt.Sprintf("load ext config failed: %s", err))
	}
}

func ExtConf() *ExtConfig {
	return extConfig
}
