package routes

import (
	"artiz-pogo-api/controllers"
	"artiz-pogo-api/middlewares"

	"github.com/go-chi/chi"
)

func PlaylistRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/playlists", func(r chi.Router) {
		r.Use(middlewares.AuthCtx)

		r.Get("/", controllers.GetAllPlaylists)
		r.Get("/{playlistID}", controllers.GetPlaylist)
		r.Post("/", controllers.CreatePlaylist)
		r.Put("/{playlistID}", controllers.UpdatePlaylistDetails)
	})

	return r
}
