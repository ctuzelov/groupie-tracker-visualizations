package main

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func ArtistsPage(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url[2])
	if r.URL.Path != "/artist/"+url[2] {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	if id > 52 || id < 1 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	ts, err := template.ParseFiles("./ui/html/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = UnmarshallArtists()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = UnmarshallRelations()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, Artists[id-1])
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
