package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	g.l.Println("Good bye")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(resp, "request failed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(resp, "hello %s", data)
}
