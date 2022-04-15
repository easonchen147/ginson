package main

import (
	"ginson/api"
	"github.com/easonchen147/foundation"
)

func main() {
	foundation.StartServer(api.RegisterRoutes)
}