package main

import (
	"ginson/api"
	"ginson/foundation"
)

func main() {
	foundation.StartServer(api.RegisterRoutes)
}