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
	http.ListenAndServe(":8081", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello friend")
	myTemplates.ExecuteTemplate(w, "index", nil)
}
