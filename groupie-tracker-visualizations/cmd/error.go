package main

import (
	"net/http"
	"text/template"
)

type ErrorMessage struct {
	Status       int
	ErrorMessage string
}

func ErrorHandler(w http.ResponseWriter, status int) {
	ts, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	ErrMessage := ErrorMessage{Status: status, ErrorMessage: http.StatusText(status)}
	err = ts.Execute(w, ErrMessage)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
