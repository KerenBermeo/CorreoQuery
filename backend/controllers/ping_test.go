package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KerenBermeo/CorreoQuery/controllers"
)

func TestPingHandler(t *testing.T) {
	// Crear un Request HTTP de prueba utilizando httptest.NewRequest()
	req := httptest.NewRequest("GET", "/ping", nil)
	// Crear un ResponseRecorder para grabar la respuesta
	w := httptest.NewRecorder()

	// Llamar a la función que estamos probando
	controllers.Ping(w, req)

	// Verificar la respuesta
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Verificar el cuerpo de la respuesta
	expectedBody := "Pong!"
	if body := w.Body.String(); body != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, body)
	}
}
