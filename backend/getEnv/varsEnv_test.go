package getenv_test

import (
	"os"
	"testing"

	getenv "github.com/KerenBermeo/CorreoQuery/getEnv"
)

func TestSetBasicAuth(t *testing.T) {
	// Configurar variables de entorno para la prueba
	os.Setenv("ZINC_USER", "test_user")
	os.Setenv("ZINC_PASSWORD", "test_password")

	// Verificar que las variables de entorno se hayan configurado correctamente
	if user := os.Getenv("ZINC_USER"); user == "" {
		t.Error("La variable de entorno ZINC_USER está vacía")
	}

	if password := os.Getenv("ZINC_PASSWORD"); password == "" {
		t.Error("La variable de entorno ZINC_PASSWORD está vacía")
	}
}

func TestGetZincSearchServerURL(t *testing.T) {
	// Configurar una variable de entorno para la prueba
	os.Setenv("ZINC_SERVER_HOST", "test_server_url")

	// Verificar que la función GetZincSearchServerURL devuelva un valor no vacío
	if url := getenv.GetZincSearchServerURL(); url == "" {
		t.Error("La URL del servidor de búsqueda de Zinc está vacía")
	}
}

func TestGetNameIndex(t *testing.T) {
	// Configurar una variable de entorno para la prueba
	os.Setenv("ZINC_SERVER_NAME_INDEX", "test_name_index")

	// Verificar que la función GetNameIndex devuelva un valor no vacío
	if index := getenv.GetNameIndex(); index == "" {
		t.Error("El índice del nombre del servidor de Zinc está vacío")
	}
}
