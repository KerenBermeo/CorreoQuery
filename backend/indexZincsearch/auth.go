package indexzincsearch

import (
	"net/http"
	"strings"
)

func MakeRequestWithAuth(method string, url string, body string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	// zincUser := os.Getenv("ZINC_FIRST_ADMIN_USER")
	// zincPass := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")
	zincUser := "admin"
	zincPass := "Complexpass#123"
	req.SetBasicAuth(zincUser, zincPass)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
