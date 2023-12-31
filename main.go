package main

import (
	"ginson/handler"
	"github.com/easonchen147/foundation"
	"github.com/easonchen147/foundation/util"
)

func main() {
	util.InitGoPool(1000) // init goroutine pool, max 1000 size
	foundation.StartServer(handler.RegisterRoutes)
}
