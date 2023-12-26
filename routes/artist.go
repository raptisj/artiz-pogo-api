package routes

import (
	"artiz-pogo-api/controllers"

	"github.com/go-chi/chi"
)

func ArtistRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/artists", controllers.GetArtists)
	r.Route("/artists/{artistID}", func(r chi.Router) {
		r.Get("/", controllers.GetSingleArtist)
	})

	return r
}
