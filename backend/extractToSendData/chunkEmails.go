package data

import "runtime"

// ChunkEmails divide una lista de rutas de correos electrónicos en fragmentos más pequeños.
func ChunkEmails(mailsPaths []string) [][]string {
	var chunks [][]string
	numCPU := runtime.NumCPU()
	numMails := len(mailsPaths)

	chunkSize := (numMails + numCPU - 1) / numCPU

	for i := 0; i < numMails; i += chunkSize {
		end := i + chunkSize
		if end > numMails {
			end = numMails
		}

		chunks = append(chunks, mailsPaths[i:end])
	}
	return chunks
}
