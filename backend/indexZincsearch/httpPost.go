package indexzincsearch

import (
	"io"
	"log"
	"net/http"
)

func httpPOST(url string, body []byte) {
	strBody := string(body)
	req, err := MakeRequestWithAuth("POST", url, strBody)
	if err != nil {
		log.Print(err)
		return
	}

	log.Printf("Posting to: %s...", url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return
	}
	defer res.Body.Close()

	log.Printf("Zinc server response code: %d", res.StatusCode)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("Zinc server response body: %s", string(resBody))
}
