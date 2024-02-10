package getenv

import "net/http"

func SetBasicAuth(req *http.Request) {
	//username := os.Getenv("ZINC_USER")
	//password := os.Getenv("ZINC_PASSWORD")
	//req.SetBasicAuth(username, password)
	req.SetBasicAuth("admin", "Complexpass#123")
}
