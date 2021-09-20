package handlers

import (
	"hl/web/data"
	"log"
	"net/http"
)

type productsController struct {
	*log.Logger
}

func NewProductsController(logger *log.Logger) *productsController {
	return &productsController{Logger: logger}
}

func (p productsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w)
	case http.MethodPut:
		p.updateProduct(w, r)
	case http.MethodPost:
		p.addProduct(w, r)
	}
}

func (p productsController) getProducts(rw http.ResponseWriter) {
	p.Println("Getting products")

	pl := data.GetProducts()

	err := pl.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusInternalServerError)
	}
}

func (p productsController) updateProduct(rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)

	p.Println("Updating product: ", product.ID)

	if ok := data.UpdateProduct(product); ok {
		p.Println("Product updated!")
	}

	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p productsController) addProduct(rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)

	p.Println("Adding product: ", product.ID)

	if id, ok := data.AddProduct(product); ok {
		p.Println("Product Added. Id: ", id)
	}

	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
