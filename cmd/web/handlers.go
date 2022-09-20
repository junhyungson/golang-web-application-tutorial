package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	// check request url path as "/" else throw an 404 error
	// else any other paths that matches "/" will succeeds
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	// w.Write([]byte("Hello from Snippet"))
}

// show snippet
// ANY
func snippetView(w http.ResponseWriter, r *http.Request) {
	// /snippet/view

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 { //
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "display a specific snippet with id %d", id)
}

// create a snippet
// POST
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// /snippet/create

	if r.Method != http.MethodPost {
		// make sure this handler only handles a post method
		// anything else will result in not allowed
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// http.Error(w, "Method Not Allowd", 405)  // better way to send status code error with a message
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // instead of hardcode, use constants
		return
	}
	w.Write([]byte("creating a new snipasdfasdfpet"))

}
