package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {

	files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))

	fmt.Fprintf(writer, "Hello, world\n")
}
