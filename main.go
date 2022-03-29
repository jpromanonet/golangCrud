package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {|
	http.HandleFunc("/", Index)
	log.Println("Server is up and running")
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello friend")
}
