package data

import (
	"fmt"

	zinc "github.com/KerenBermeo/CorreoQuery/indexZincsearch"
)

func SendToZincSearch(data []string) {

	var format string

	for _, items := range data {
		format = items
	}

	fmt.Println(format)
	// url para operación masiva
	//url := "http://localhost:4080/api/_bulkv2"
	url := "http://localhost:4080/api/_bulk"

	// Enviar datos al servidor
	zinc.HttpPOST(url, format)
}
