package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

// Función para manejar errores al llamar a wrapper.GetIndexNamesList
func handleWrapperError(err error, w http.ResponseWriter, r *http.Request) {
	// Loguea el error para propósitos de depuración
	log.Printf("Error al obtener los nombres de índice: %v", err)

	// Define el mensaje de error para el cliente
	errorMsg := "Error al obtener los nombres de índice"

	// Determina el código de estado HTTP apropiado en función del tipo de error
	statusCode := http.StatusInternalServerError
	if strings.Contains(err.Error(), "Unauthorized") {
		statusCode = http.StatusUnauthorized
		errorMsg = "No autorizado para acceder a los nombres de índice"
	} else if strings.Contains(err.Error(), "NotFound") {
		statusCode = http.StatusNotFound
		errorMsg = "El recurso solicitado no se encontró"
	}

	// Establece el código de estado antes de llamar a render.JSON
	w.WriteHeader(statusCode)
	// Renderiza el mensaje de error JSON en la respuesta HTTP
	render.JSON(w, r, map[string]string{"error": errorMsg})
}
