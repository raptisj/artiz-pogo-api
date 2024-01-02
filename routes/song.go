package routes

import (
	"artiz-pogo-api/controllers"
	"artiz-pogo-api/middlewares"

	"github.com/go-chi/chi"
)

func SongRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", controllers.GetAllSongsFromArtist)
		r.Get("/{songID}", controllers.GetSingleSong)

	})

	r.Route("/songs/{songID}", func(r chi.Router) {
		r.Use(middlewares.AuthCtx)

		r.Post("/likes/add", controllers.AddLikeToSong)
		r.Delete("/likes/remove", controllers.RemoveLikeFromSong)
	})

	return r
}
