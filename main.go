package main

import (
	"artiz-pogo-api/models"
	"artiz-pogo-api/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// stop processing after a minute
	r.Use(middleware.Timeout(60 * time.Second))

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := models.Config{
		Host:     os.Getenv("DB_DOCKER_HOST"), // DB_DOCKER_HOST has a value of db. It's a workaround so it can play nice with docker
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	r.Mount("/artists", routes.ArtistRoutes())
	r.Mount("/songs", routes.SongRoutes())
	r.Mount("/users", routes.AuthRoutes())
	r.Mount("/playlists", routes.PlaylistRoutes())

	http.ListenAndServe(":8080", r)
}
