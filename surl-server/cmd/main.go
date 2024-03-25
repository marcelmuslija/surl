package main

import (
	"fmt"
	"net/http"
	"surl-server/internal/config"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", healthcheck)

	config := config.Get()

	fmt.Println("Server started on port", config.Port)
	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		fmt.Println("Server failed to start.")
	}
}
