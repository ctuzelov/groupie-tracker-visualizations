package main

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = UnmarshallArtists()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, Artists)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
	}
}
