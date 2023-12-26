package controllers

import (
	"artiz-pogo-api/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	db := models.DB
	artists := []models.Artist{}

	db.Find(&artists)
	jsonData, err := json.Marshal(artists)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetSingleArtist(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "artistID")

	db := models.DB
	artist := models.Artist{ID: artistID}

	db.Find(&artist)
	jsonData, err := json.Marshal(artist)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// artist := models.Artist{
// 	Name:       "Jimi Hendrix",
// 	BirthYear:  1942,
// 	DeathYear:  1970,
// 	Instrument: "Guitarist, singer",
// 	Genre:      "Rock, funk, psychedelic rock",
// 	Bio:        "One of the most innovative and influential guitarists of all time, known for his electrifying stage presence and groundbreaking use of distortion and feedback.",
// }
