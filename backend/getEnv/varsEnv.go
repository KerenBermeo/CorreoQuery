package getenv

import (
	"fmt"
	"net/http"
	"os"
)

func SetBasicAuth(req *http.Request) error {
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	if username == "" {
		return fmt.Errorf("ZINC_FIRST_ADMIN_USER")
	}

	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")
	if password == "" {
		return fmt.Errorf("ZINC_FIRST_ADMIN_USER no está configurado")
	}

	req.SetBasicAuth(username, password)
	return nil
}

func GetZincSearchServerURL() string {
	return os.Getenv("ZINC_SEARCH_URL")

}

func GetNameIndex() string {
	return os.Getenv("ZINC_SERVER_NAME_INDEX")
}

func GetRootDirectory() string {
	return os.Getenv("BACK_ROOT_DIRECTORY")
}
