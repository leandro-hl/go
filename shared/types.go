package shared

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

//any pointer of type Page could run this function
func (p *Page) Save() error {
	filename := "data/" + p.Title + ".txt"

	//read-write permissions for current user only
	return ioutil.WriteFile(filename, p.Body, 0600)
}
