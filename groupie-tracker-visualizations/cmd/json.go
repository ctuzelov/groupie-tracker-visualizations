package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Artist struct {
	ID            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  int                 `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	Locations     string              `json:"locations"`
	ConcertDates  string              `json:"concertDates"`
	DatesLocation map[string][]string `json:"datesLocations"`
	Relations     string              `json:"relations"`
}

type Relations struct {
	Index []struct {
		ID            int                 `json:"id"`
		DatesLocation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

var (
	Artists  []Artist
	Relation Relations
)

func UnmarshallArtists() error {
	if len(Artists) != 0 {
		return nil
	}
	url := "https://groupietrackers.herokuapp.com/api/artists"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	jsonErr := json.Unmarshal(body, &Artists)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func UnmarshallRelations() error {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	jsonErr := json.Unmarshal(body, &Relation)
	if jsonErr != nil {
		log.Println("jsonErr")
		return jsonErr
	}
	for i := range Artists {
		Artists[i].DatesLocation = Relation.Index[i].DatesLocation
		FormatDates(i)
	}
	return nil
}

func FormatDates(i int) {
	res := make(map[string][]string)
	for key := range Artists[i].DatesLocation {
		res[strings.Title(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(key, "_", " "), "-", ", "), " usa", " USA"), " uk", " UK"))] = Artists[i].DatesLocation[key]
		delete(Artists[i].DatesLocation, key)
	}
	Artists[i].DatesLocation = res
}
