package router

import (
	"log"

	c "github.com/KerenBermeo/CorreoQuery/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", c.Ping)
		r.Get("/indexes", c.GetIndexNamesList)
		r.Post("/emails", c.GetEmails)
		r.Post("/makesearch", c.MakeSearch)
	})

	log.Println("Routes initialized successfully")
}
