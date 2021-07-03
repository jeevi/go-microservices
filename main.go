package main

import (
	"go/microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()
	http.ListenAndServe(":8080", nil)
}
