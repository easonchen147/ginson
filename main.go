package main

import (
	"ginson/handler"
	"github.com/easonchen147/foundation"
)

func main() {
	foundation.StartServer(handler.RegisterRoutes)
}