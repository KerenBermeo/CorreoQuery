package data

import (
	"bufio"
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
				continue
			}
			headerTextBuffer.WriteString(line + "\n")
		}

		for scanner.Scan() {
			bodyTextBuffer.WriteString(scanner.Text() + "\n")
		}

		parsedEmail, _ := ParseEmailWithHeadersAndBody(headerTextBuffer.String(), bodyTextBuffer.String())

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
