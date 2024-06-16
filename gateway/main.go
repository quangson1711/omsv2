package main

import (
	"common"
	_ "github.com/joho/godotenv/autoload"
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
