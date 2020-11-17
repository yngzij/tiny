package main

import (
	"tidy/service"

)

func main() {
	server := service.NewEchoServer()
	server.Run()
}
