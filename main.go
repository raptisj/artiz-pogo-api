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

// type Artist struct {
// 	Name       string `json:"name"`
// 	BirthYear  string `json:"birth_year"`
// 	DeathYear  string `json:"death_year"`
// 	Instrument string `json:"instrument"`
// 	Genre      string `json:"genre"`
// 	Bio        string `json:"bio"`
// }

// var Artists []*Artist

// func getArtists(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(Artists)
// }

//	func getSingleArtist(w http.ResponseWriter, r *http.Request) {
//		artist := new(Artist)
//		if err := json.NewDecoder(r.Body).Decode(artist); err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			w.Write([]byte("Please enter a correct Todo!!"))
//			return
//		}
//		Artists = append(Artists, artist)
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte("Todo added!!"))
//	}
// func GetSingleArtist(w http.ResponseWriter, r *http.Request) {
// 	artistID := chi.URLParam(r, "artistID")
// 	fmt.Println(artistID)
// 	fmt.Println("====")

// 	db := models.DB
// 	artist := models.Artist{ID: artistID}

// 	db.Find(&artist)
// 	jsonData, err := json.Marshal(artist)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonData)
// }

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// stop processing after a minute
	r.Use(middleware.Timeout(60 * time.Second))

	// r.Route("/artists", func(r chi.Router) {
	// 	r.Get("/", getArtists)
	// 	r.Get("/{artistID}", getSingleArtist)
	// })

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	r.Mount("/", routes.Routes())

	http.ListenAndServe(":4000", r)
}

// func Routes() chi.Router {
// 	r := chi.NewRouter()

// 	r.Get("/", controllers.GetArtists)
// 	r.Route("/{artistID}", func(r chi.Router) {
// 		r.Get("/", controllers.GetSingleArtist)
// 	})

// 	return r
// }

// func ArtistRoutes() chi.Router {
// 	r := chi.NewRouter()
// 	// bookHandler := BookHandler{
// 	// 	storage: BookStore{},
// 	// }
// 	r.Get("/", getArtists)
// 	r.Get("/{artistID}", getSingleArtist)
// 	// r.Post("/", bookHandler.CreateBook)
// 	// r.Get("/{id}", bookHandler.GetBooks)
// 	// r.Put("/{id}", bookHandler.UpdateBook)
// 	// r.Delete("/{id}", bookHandler.DeleteBook)
// 	return r
// }
