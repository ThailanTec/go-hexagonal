package main

import (
	"fmt"

	server2 "github.com/ThailanTec/go-hexagonal/adapter/web/server"
)

func main() {
	server := server2.NewWebServer()

	server.Serve()
	fmt.Println("Servidor rolando")
}
