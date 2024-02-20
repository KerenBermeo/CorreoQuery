package data

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/KerenBermeo/CorreoQuery/model"
)

func ParseEmailFromPath(paths []string, maxBatch int, wg *sync.WaitGroup) {

	var email []model.Email

	for i, path := range paths {
		fd, err := os.Open(path)
		if err != nil {
			continue
		}
		defer fd.Close()

		scanner := bufio.NewScanner(fd)

		var headerTextBuffer strings.Builder
		var bodyTextBuffer strings.Builder

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			if len(line) > bufio.MaxScanTokenSize {
				// Si la línea es demasiado larga, pasa a la siguiente.
				continue
			}
			headerTextBuffer.WriteString(line + "\n")
		}

		// Leer el cuerpo del correo electrónico
		for scanner.Scan() {
			bodyTextBuffer.WriteString(scanner.Text() + "\n")
		}

		parsedEmail, _ := ParseEmail(headerTextBuffer.String(), bodyTextBuffer.String())

		email = append(email, parsedEmail)

		batchSize := len(paths)

		if (i+1)%maxBatch == 0 {
			log.Printf("Uploading bulk %v / %v", i+1, batchSize)
			ConcurrentParsedEmailJson(email)
			email = nil
		} else if email != nil && (i+1) == batchSize {
			log.Printf("Uploading bulk %v / %v", i+1, batchSize)
			ConcurrentParsedEmailJson(email)
		}

	}

	wg.Done()
}

func ParseEmail(headerText string, bodyText string) (model.Email, error) {

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
		if shouldExclude(line) {
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

func shouldExclude(line string) bool {
	excludedHeaders := []string{
		"Mime-Version: ",
		"Content-Type: ",
		"Content-Transfer-Encoding: ",
		"X-From: ",
		"X-To: ",
		"X-cc: ",
		"X-bcc: ",
		"X-Folder: ",
		"X-Origin: ",
		"X-FileName: ",
	}
	for _, excluded := range excludedHeaders {
		if strings.HasPrefix(line, excluded) {
			return true
		}
	}
	return false
}
