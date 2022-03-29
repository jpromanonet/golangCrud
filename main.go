package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Let's create some templates
var myTemplates = template.Must(template.ParseGlob("templates/*"))

func main() {
	// We call the index function
	http.HandleFunc("/", Index)
	// Show in console that the server is up and running
	fmt.Println("Server is up and running")
	// We serve in the following port
	http.ListenAndServe(":8082", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// let's call the templates
	myTemplates.ExecuteTemplate(w, "header", nil)
	myTemplates.ExecuteTemplate(w, "index", nil)
	myTemplates.ExecuteTemplate(w, "footer", nil)
}
