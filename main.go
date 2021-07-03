package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jeevi/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()
	http.ListenAndServe(":8080", nil)
}
