package getenv

import "net/http"

func SetBasicAuth(req *http.Request) {
	//username := os.Getenv("ZINC_USER")
	//password := os.Getenv("ZINC_PASSWORD")
	//req.SetBasicAuth(username, password)
	req.SetBasicAuth("admin", "Complexpass#123")
}

func GetZincSearchServerURL() string {
	// return os.Getenv("ZINC_SERVER_HOST")
	return "http://localhost:4080/"
}

func GetNameIndex() string {
	//return os.Getenv("ZINC_SERVER_NAME_INDEX")
	return "email"
}
