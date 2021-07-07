package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jeevi/go-microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(resp, req)
		return
	}

	if req.Method == http.MethodPut {
		path := req.URL.Path
		reg := regexp.MustCompile(`/([0-9]+)`)
		matches := reg.FindAllStringSubmatch(path, -1)

		if len(matches) != 1 {
			http.Error(resp, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(matches[0]) != 2 {
			http.Error(resp, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := matches[0][1]
		id, _ := strconv.Atoi(idString)

		p.l.Println(id)
		p.updateProducts(id, resp, req)
		return
	}

	if req.Method == http.MethodPost {
		p.addProduct(resp, req)
		return
	}

	resp.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(resp http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET")
	lp := data.GetProducts()
	err := lp.ToJSON(resp)

	if err != nil {
		http.Error(resp, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) updateProducts(id int, resp http.ResponseWriter, req *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(resp, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(resp, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(resp, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(resp http.ResponseWriter, req *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)

	if err != nil {
		http.Error(resp, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}
