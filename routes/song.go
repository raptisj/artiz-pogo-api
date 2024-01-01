package routes

import (
	"artiz-pogo-api/controllers"

	"github.com/go-chi/chi"
)

func SongRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/songs", controllers.GetAllSongFromArtist)
	r.Route("/songs/{songID}", func(r chi.Router) {
		r.Get("/", controllers.GetSingleSong)
	})

	return r
}
