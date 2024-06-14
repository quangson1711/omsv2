package main

import (
	"common"
	"log"
	"net/http"
)

var (
	httpAddr = common.GetEnv("HTTP_ADDR", ":3000")
)

func main() {

	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Print("Starting server on ", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal(err)
	}
}
