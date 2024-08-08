package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/yash91989201/go_cart/configs"
	"github.com/yash91989201/go_cart/routes"
)

func main() {

	PORT := ":" + configs.GetEnv().PORT

	router := chi.NewRouter()
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           3000,
			},
		),
	)

	db := configs.GetDB()

	v1Router := chi.NewRouter()
	routes.Register(v1Router, db)

	router.Mount("/api/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    PORT,
	}

	log.Printf("Starting server on port %s", PORT)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
