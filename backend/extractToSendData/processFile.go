package data

import (
	"bufio"
	"os"
	"strings"

	"github.com/KerenBermeo/CorreoQuery/model"
)

func ProcessFile(file *os.File) model.Email {

	scanner := bufio.NewScanner(file)

	var bodyText strings.Builder

	header := make(map[string]string)

	foundEmptyLine := false

	var bodyTextBuffer strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if len(line) > bufio.MaxScanTokenSize {
			continue
		}

		if foundEmptyLine {
			bodyText.WriteString(line) // Agrega la línea al cuerpo del texto
			continue
		}

		pass := shouldExclude(line)

		if pass {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				header[key] = value
			}
		}
	}
	for scanner.Scan() {
		bodyTextBuffer.WriteString(scanner.Text() + "\n")
	}

	email := model.Email{
		From:    header["From"],
		To:      strings.Split(header["To"], ","),
		Subject: header["Subject"],
		Date:    header["Date"],
		Content: bodyTextBuffer.String(),
	}

	return email

}
