package data

import (
	"encoding/json"
	"fmt"
	hlu "hl/web/utils"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) IsLookup(key interface{}) bool {
	return p.ID == key.(int)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)

	return e.Decode(p)
}

type Products []*Product

func (p Products) Len() int {
	return len(p)
}

func (p Products) Get(i int) hlu.Searchable {
	return p[i]
}

type ByIDDescending Products

func (p ByIDDescending) Len() int {
	return len(p)
}

func (p ByIDDescending) Less(i, j int) bool {
	return p[i].ID < p[j].ID
}

func (p ByIDDescending) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Products) ToJSON(w io.Writer) error {
	// Encoder that writes to w
	e := json.NewEncoder(w)
	fmt.Println()
	return e.Encode(p)
}

func (p Products) GetByID(id int) (int, *Product) {
	index, result := hlu.LinearSearch(p, id)
	prod := result.(*Product)
	return index, prod
}

func (p Products) Update(product *Product) bool {
	index, _ := p.GetByID(product.ID)

	p[index] = product

	return true
}

func GetProducts() Products {
	return products
}

func AddProduct(p *Product) (int, bool) {
	p.ID = nextIndex()
	products = append(products, p)
	return p.ID, true
}

func UpdateProduct(p *Product) bool {
	return products.Update(p)
}

func init() {
	pl := []*Product{
		{
			Name:        "Latte",
			Description: "Frothy milky coffee",
			Price:       2.45,
			SKU:         "abc323",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
		{
			Name:        "Espresso",
			Description: "Short and strong coffee without milk",
			Price:       1.99,
			SKU:         "fjd34",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
	}

	for i := range pl {
		pl[i].ID = nextIndex()
	}

	products = pl
}

func nextIndex() int {
	indexer++
	return indexer
}

var indexer = 1
var products Products
