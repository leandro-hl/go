package main

import (
	"fmt"
	"hl/shared"
)

func main() {
	title := "Test"
	page := &shared.Page{Title: title, Body: []byte("This is a test page")}

	page.Save()

	read, _ := shared.LoadPage(title)

	fmt.Println(string(read.Body))
}
