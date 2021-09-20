package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hello struct {
	*log.Logger
}

func (h hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.Println("Handle hello request")

	// read body
	b, err := ioutil.ReadAll(request.Body)

	if err != nil {
		h.Println("Found error during body reading...")

		http.Error(writer, "Unable to read request body", http.StatusBadRequest)
	}

	//write response
	fmt.Fprintf(writer, "Hello %s", b)
}

func NewHello(l *log.Logger) http.Handler {
	return &hello{l}
}
