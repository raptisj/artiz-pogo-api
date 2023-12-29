package routes

import (
	"artiz-pogo-api/controllers"

	"github.com/go-chi/chi"
)

func AuthRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/users/signup", controllers.Signup)
	r.Post("/users/login", controllers.Login)
	r.Get("/users/current", controllers.CurrentUser)

	return r
}
