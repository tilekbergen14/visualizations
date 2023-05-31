package controllers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	num , err := strconv.Atoi(id)
	if err != nil || (num > 52 || num < 1) {
		errorHandler(w, 400)
		return
	}
	artist := getArtist(id)
	date := getDate(artist.ConcertDates)
	location := getLocation(artist.Locations)
	relation := getRelation(artist.Relations)
	t, err := template.ParseFiles("views/" + "artist" + ".html")
	if err != nil {
		return
	}
	t.Execute(w, models.Data{Artist: artist, Date: date, Location: location, Relation: relation})
}

func getArtist(id string) models.Artist {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	artistsData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist models.Artist
	json.Unmarshal(artistsData, &artist)
	
	return artist
}

func getDate(link string) models.Date {
	res, err := http.Get(link)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var date models.Date
	json.Unmarshal(data, &date)
	return date
}

func getLocation(link string) models.Location {
	res, err := http.Get(link)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var location models.Location
	json.Unmarshal(data, &location)
	return location
}

func getRelation(link string) models.Relation {
	res, err := http.Get(link)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var relation models.Relation
	json.Unmarshal(data, &relation)
	return relation
}
