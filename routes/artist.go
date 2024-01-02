package routes

import (
	"artiz-pogo-api/controllers"
	"artiz-pogo-api/middlewares"

	"github.com/go-chi/chi"
)

func ArtistRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/artists", controllers.GetArtists)
	r.Get("/artists/{artistID}", controllers.GetSingleArtist)

	r.Route("/artists/{artistID}/follow", func(r chi.Router) {
		r.Use(middlewares.AuthCtx)

		r.Post("/add", controllers.FollowArtist)
		r.Delete("/remove", controllers.UnfollowArtist)
	})

	return r
}
