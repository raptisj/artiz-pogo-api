package controllers

import (
	"artiz-pogo-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetAllPlaylists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)

	db := models.DB
	playlists := []models.Playlist{}

	db.Where("user_id", userId).Find(&playlists)
	jsonData, err := json.Marshal(playlists)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistID := chi.URLParam(r, "playlistID")
	playlistIDUint, err := strconv.ParseUint(playlistID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db := models.DB
	playlist := models.Playlist{ID: uint(playlistIDUint)}

	db.Find(&playlist)
	jsonData, err := json.Marshal(playlist)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)

	var playlistBody models.Playlist

	if err := json.NewDecoder(r.Body).Decode(&playlistBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide the correct input!!"))
		return
	}

	db := models.DB
	playlist := models.Playlist{
		UserID:    userId,
		Title:     playlistBody.Title,
		SongCount: 0,
		SongIds:   []int{},
	}

	db.Create(&playlist)

	fmt.Fprintf(w, "Playlist created successfully")
	w.WriteHeader(http.StatusOK)
}

func UpdatePlaylistDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)

	playlistID := chi.URLParam(r, "playlistID")
	playlistIDUint, err := strconv.ParseUint(playlistID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	var playlistBody models.Playlist

	if err := json.NewDecoder(r.Body).Decode(&playlistBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide the correct input!!"))
		return
	}

	db := models.DB
	playlist := models.Playlist{}

	models.DB.Where("id = ? AND user_id = ?", uint(playlistIDUint), userId).First(&playlist)

	if playlistBody.Title != "" {
		playlist.Title = playlistBody.Title
	}

	if playlistBody.Description != "" {
		playlist.Description = playlistBody.Description
	}

	db.Save(&playlist)

	fmt.Fprintf(w, "Playlist updated successfully")
	w.WriteHeader(http.StatusOK)
}
