package shared

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func LoadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	//returns the pointer to the created Page object
	//so there's no need to copy it again externally (duplicating memory usage)
	return &Page{Title: title, Body: body}, nil
}

func LoadPage2() {

}

//as it only has an int, why making an struct?
type Counter struct {
	count int
}

//No need for an struct
type CounterB int

func (c *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.count++

	fmt.Fprintf(w, "request made %d times\n", c.count)
}

func (c *CounterB) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*c++

	fmt.Fprintf(w, "request made %d times\n", *c)
}

//Notify a page has been visited
//A channel that sends a notification on each visit.
//Probably it should be buffered.
type Channel chan *http.Request

func (c Channel) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c <- req

	fmt.Fprintf(w, "In code we trust.")
}

//Adding a function as a receiver.
type Args func(http.ResponseWriter, *http.Request)

//function of type Args
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, os.Args)
}

func (fn Args) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fn(w, req)
}
