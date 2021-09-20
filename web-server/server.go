package main

import (
	"fmt"
	"hl/shared"
	"hl/web/utils"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	pointerToPageObjectCreated, err := shared.LoadPage(title)

	utils.ExecOrExecError(
		err,
		func() {
			http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		},
		func() {
			renderTemplate(w, "view", pointerToPageObjectCreated)
		},
	)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	pointerToPageObjectCreated, err := shared.LoadPage(title)

	utils.ExecOnError(err, func() { pointerToPageObjectCreated = &shared.Page{Title: title} })

	renderTemplate(w, "edit", pointerToPageObjectCreated)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	//creates de object ones
	pointerToPageObject := &shared.Page{Title: title, Body: []byte(body)}

	//creates de obj twice, created and then copied to variable.
	//pointerToPageObject := shared.Page{Title: title, Body: []byte(body)}

	utils.ExecOrError(
		w,
		pointerToPageObject.Save(),
		func() {
			http.Redirect(w, r, "/view/"+title, http.StatusFound)
		},
	)
}

var templates = utils.GetTemplates()

func renderTemplate(w http.ResponseWriter, tmpl string, p *shared.Page) {
	utils.SendError(w, templates().ExecuteTemplate(w, tmpl+".html", p))
}

func testInterface() {
	command := utils.SelectCommand("write")

	command.Exec()
}

func main() {
	port := "8080"

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", utils.MakeHandler(viewHandler))
	http.HandleFunc("/edit/", utils.MakeHandler(editHandler))
	http.HandleFunc("/save/", utils.MakeHandler(saveHandler))

	testInterface()
	fmt.Printf("listening on port %v...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
