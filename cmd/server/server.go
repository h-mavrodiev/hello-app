package main

import (
	"sync"

	server "github.com/h-mavrodiev/hello-app/pkg/server"
)

func main() {
	var m sync.Mutex
	webAppPort := server.GetEnv("WEBAPP_PORT", "8080")
	r := server.HelloAppRouter(&m)
	r.Run(":" + webAppPort)
}
