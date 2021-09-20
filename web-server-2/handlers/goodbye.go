package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type bye struct {
	*log.Logger
}

func NewBye(logger *log.Logger) *bye {
	return &bye{Logger: logger}
}

func (b bye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	b.Println("Handle goodbye request")

	//write response
	fmt.Fprint(writer, "Good bye")
}
