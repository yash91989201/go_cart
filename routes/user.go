package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/yash91989201/go_cart/controllers"
	"github.com/yash91989201/go_cart/internal/database"
)

func RegisterUserRoutes(router chi.Router, db *database.Queries) {
	userController := controllers.UserControllers{
		DB: db,
	}

	router.Group(func(route chi.Router) {
		route.Post("/sign-up", userController.CreateUser)
		route.Post("/login", userController.LoginUser)
	})
}
