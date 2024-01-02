package routes

import (
	"artiz-pogo-api/controllers"
	"artiz-pogo-api/middlewares"

	"github.com/go-chi/chi"
)

func SongRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/songs", controllers.GetAllSongsFromArtist)
	r.Get("/songs/{songID}", controllers.GetSingleSong)

	r.Route("/songs/{songID}/likes", func(r chi.Router) {
		r.Use(middlewares.AuthCtx)

		r.Post("/add", controllers.AddLikeToSong)
		r.Delete("/remove", controllers.RemoveLikeFromSong)
	})

	return r
}
