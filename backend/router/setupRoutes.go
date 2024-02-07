package router

import (
	"github.com/KerenBermeo/CorreoQuery/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Get("/api/index", controllers.ListIndexes)
	//r.Get("/api/index_name", controllers.ListIndexes)
}
