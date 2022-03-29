package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Solicitud para acceder a funcion Index
	http.HandleFunc("/", Index)
	// Mensaje en consola
	log.Println("Server is up and running")
	// Puerto en el que se sirve la app
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello friend")
}
