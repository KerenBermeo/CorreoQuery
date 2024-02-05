package indexzincsearch

import (
	"bytes"
	"net/http"
)

func MakeRequestWithAuth(method string, url string, data []byte) (*http.Request, error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	zincUser := "admin"
	zincPass := "Complexpass#123"
	req.SetBasicAuth(zincUser, zincPass)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
