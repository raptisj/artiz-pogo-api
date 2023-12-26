package controllers

import (
	"artiz-pogo-api/models"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func GetAllSongById(w http.ResponseWriter, r *http.Request) {
	db := models.DB
	songs := []models.Song{}
	songIDs := r.URL.Query().Get("songIDs")
	songIDsArray := strings.Split(songIDs, ",")

	db.Find(&songs, songIDsArray)
	jsonData, err := json.Marshal(songs)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetSingleSong(w http.ResponseWriter, r *http.Request) {
	songID := chi.URLParam(r, "songID")

	db := models.DB
	song := models.Song{ID: songID}

	db.Find(&song)
	jsonData, err := json.Marshal(song)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
