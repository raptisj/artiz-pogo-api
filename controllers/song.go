package controllers

import (
	"artiz-pogo-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetAllSongFromArtist(w http.ResponseWriter, r *http.Request) {
	db := models.DB
	songs := []models.Song{}
	artistID := r.URL.Query().Get("artistID")
	artistIDUint, err := strconv.ParseUint(artistID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db.Where("artist_id", artistIDUint).Find(&songs)
	jsonData, err := json.Marshal(songs)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetSingleSong(w http.ResponseWriter, r *http.Request) {
	songID := chi.URLParam(r, "songID")
	songIDUint, err := strconv.ParseUint(songID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db := models.DB
	song := models.Song{ID: uint(songIDUint)}

	db.Find(&song)
	jsonData, err := json.Marshal(song)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
