package main

import (
	"fmt"
	"net/http"
	"surl-server/internal/server"
)

func main() {
	server, err := server.New()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server started on port", server.Config.Port)
	if err := http.ListenAndServe(server.Config.Port, server.Router); err != nil {
		fmt.Println("Server failed to start.")
	}
}
