package main

import (
	"fmt"
	"ginson/pkg/conf"
	"ginson/platform"
)

func main() {
	// 初始化相关依赖组件
	err := platform.Initialize(conf.AppConf)
	if err != nil {
		panic(fmt.Sprintf("Initialize failed: %s", err))
	}

	// 启动Web服务
	err = platform.StartServer(conf.AppConf)
	if err != nil {
		panic(fmt.Sprintf("Server started failed: %s", err))
	}
}
