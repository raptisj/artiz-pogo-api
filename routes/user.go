package routes

import (
	"artiz-pogo-api/controllers"
	"artiz-pogo-api/middlewares"

	"github.com/go-chi/chi"
)

func AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/users/signup", controllers.Signup)
	r.Post("/users/login", controllers.Login)

	r.Route("/users", func(r chi.Router) {
		r.Use(middlewares.AuthCtx)

		r.Get("/current", controllers.CurrentUser)
		r.Get("/likes/list", controllers.GetAllLikedSongs)
		r.Get("/follow/list", controllers.GetAllFollowedArtist)
	})

	return r
}
