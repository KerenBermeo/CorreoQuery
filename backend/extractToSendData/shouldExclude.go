package data

import "strings"

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
			return false
		}
	}
	return true
}
