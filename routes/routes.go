package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yash91989201/go_cart/internal/database"
	"github.com/yash91989201/go_cart/utils"
)

func Register(router chi.Router, db *database.Queries) {
	router.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJson(w, 200, "Api working")
	})

	RegisterUserRoutes(router, db)
}
