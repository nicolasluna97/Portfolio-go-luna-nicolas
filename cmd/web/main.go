package main

import (
	"log"
	"net/http"

	"github.com/nicolasluna97/Portfolio-go-luna-nicolas/internal/server"
)

func main() {
	addr := ":8080"

	router := server.NewRouter()

	log.Printf("Servidor iniciado en http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
