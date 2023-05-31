package main

import (
	"fmt"
	"groupie-tracker/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/artist/", controllers.ArtistHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Link -->   " + "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
