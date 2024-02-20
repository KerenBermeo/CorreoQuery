package data

import (
	"fmt"
	"strings"

	"github.com/KerenBermeo/CorreoQuery/model"
)

func ParseEmailWithHeadersAndBody(headerText string, bodyText string) (model.Email, error) {

	if headerText == "" || bodyText == "" {
		return model.Email{}, fmt.Errorf("header or body text is empty")
	}

	email := model.Email{}

	// Parsear el encabezado
	header := make(map[string]string)
	headerLines := strings.Split(headerText, "\n")

	for _, line := range headerLines {
		if line == "" {
			break
		}
		if ShouldExclude(line) {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			header[key] = value
		} else {
			return model.Email{}, fmt.Errorf("malformed header line: %s", line)
		}
	}

	email.MessageId = header["Message-ID"]
	email.From = header["From"]
	email.To = strings.Split(header["To"], ",")
	email.Subject = header["Subject"]
	email.Date = header["Date"]

	body := ""
	bodyLines := strings.Split(bodyText, "\n")
	for _, line := range bodyLines {
		body += line + "\n"
	}
	email.Content = body

	return email, nil
}
