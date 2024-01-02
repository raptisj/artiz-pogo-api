package controllers

import (
	"artiz-pogo-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type GetArtistResponse struct {
	Artists []models.Artist `json:"artists"`
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	db := models.DB
	artists := []models.Artist{}

	db.Find(&artists)

	response := GetArtistResponse{
		Artists: artists,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

type GetSingleArtistResponse struct {
	Artist models.Artist `json:"artist"`
	Songs  []models.Song `json:"songs"`
}

func GetSingleArtist(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "artistID")
	artistIDUint, err := strconv.ParseUint(artistID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db := models.DB
	artist := models.Artist{ID: uint(artistIDUint)}
	artistSongs := []models.Song{}

	db.Find(&artist)
	db.Where("artist_id", artistIDUint).Find(&artistSongs)

	response := GetSingleArtistResponse{
		Artist: artist,
		Songs:  artistSongs,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func FollowArtist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)

	artistID := chi.URLParam(r, "artistID")
	artistIDUint, err := strconv.ParseUint(artistID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db := models.DB
	followArtist := models.FollowArtist{
		UserID:   userId,
		ArtistID: uint(artistIDUint),
	}

	db.Create(&followArtist)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Artist followed successfully")
}

func UnfollowArtist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)

	artistID := chi.URLParam(r, "artistID")
	artistIDUint, err := strconv.ParseUint(artistID, 10, 32)
	if err != nil {
		fmt.Println("could not parse string")
	}

	db := models.DB
	liked := &models.FollowArtist{}
	db.Where("artist_id = ? AND user_id = ?", artistIDUint, userId).Unscoped().Delete(&liked)

	fmt.Fprintf(w, "You no longer follow this artist")
	w.WriteHeader(http.StatusOK)
}

func GetAllFollowedArtist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("userId").(uint)
	db := models.DB
	artists := []models.Artist{}

	artistIDsArray := []uint{}
	followed := []models.FollowArtist{}
	db.Where("user_id", userId).Find(&followed)
	for _, follow := range followed {
		artistIDsArray = append(artistIDsArray, follow.ArtistID)
	}

	db.Find(&artists, artistIDsArray)
	jsonData, err := json.Marshal(artists)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
