package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, 404)
        return
    }
	artists := getAllArtists()
	t, err := template.ParseFiles("views/" + "home" + ".html")
	if err != nil {
		return
	}
	t.Execute(w, artists)
}

func getAllArtists() []models.Artist {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	artistsData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artists []models.Artist
	json.Unmarshal(artistsData, &artists)
	return artists
}



func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	_, err := template.ParseFiles("views/error.html")
	if err != nil {
		fmt.Fprintf(w, "Internal server error!")
		return
	}
	if status == http.StatusNotFound {
		renderTemplate(w, "error", models.Data{ErrorCode: status, ErrorMessage: "Page not found!"})
	}
	if status == http.StatusMethodNotAllowed {
		renderTemplate(w, "error", models.Data{ErrorCode: status, ErrorMessage: "Method not alllowed!"})
	}
	if status == http.StatusBadRequest {
		renderTemplate(w, "error", models.Data{ErrorCode: status, ErrorMessage: "Bad request!"})
	}
	if status == http.StatusInternalServerError {
		renderTemplate(w, "error", models.Data{ErrorCode: status, ErrorMessage: "Internal Server Error!"})
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data models.Data) {
	t, err := template.ParseFiles("views/" + tmpl + ".html")
	if err != nil {
		errorHandler(w, 500)
		return
	}
	t.Execute(w, data)
}