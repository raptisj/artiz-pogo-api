package controllers

import (
	"artiz-pogo-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetAllSongsFromArtist(w http.ResponseWriter, r *http.Request) {
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

func AddLikeToSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)

	songID := chi.URLParam(r, "songID")
	songIDUint, err := strconv.ParseUint(songID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db := models.DB
	likedSong := models.LikedSong{
		UserID: userId,
		SongID: uint(songIDUint),
	}

	db.Create(&likedSong)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Song added to liked list successfully")
}

func RemoveLikeFromSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)

	songID := chi.URLParam(r, "songID")
	songIDUint, err := strconv.ParseUint(songID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db := models.DB
	liked := &models.LikedSong{}
	db.Where("song_id = ? AND user_id = ?", songIDUint, userId).Unscoped().Delete(&liked)

	fmt.Fprintf(w, "Song removed from liked list")
	w.WriteHeader(http.StatusOK)
}

func GetAllLikedSongs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)
	db := models.DB
	songs := []models.Song{}

	songIDsArray := []uint{}
	liked := []models.LikedSong{}
	db.Where("user_id", userId).Find(&liked)
	for _, like := range liked {
		songIDsArray = append(songIDsArray, like.SongID)
	}

	db.Find(&songs, songIDsArray)
	jsonData, err := json.Marshal(songs)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
