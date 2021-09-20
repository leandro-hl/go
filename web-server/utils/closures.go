package utils

import "net/http"

func MakeHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
	path := GetValidPath()
	//returned function is a closure
	//(encloses values defined outside of it)
	return func(w http.ResponseWriter, r *http.Request) {
		matches := path().FindStringSubmatch(r.URL.Path)

		if matches == nil {
			http.NotFound(w, r)
			return
		}

		//enclosed by the closure
		// The title is the second subexpression.
		fn(w, r, matches[2])
	}
}
