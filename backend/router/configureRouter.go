package router

import (
	"log"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func ConfigureRouter(r chi.Router) {

	origen := getenv.GetOrigen()
	if origen == "" {
		log.Println("Empty environment variable")
	}
	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{origen},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
	}))

	// Middleware configuration
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Route configuration
	SetupRoutes(r)
}
