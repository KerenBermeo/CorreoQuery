package data

import "runtime"

func SplitIntoChunks(paths []string) [][]string {

	var chunks [][]string
	numCPU := runtime.NumCPU()
	numMails := len(paths)

	chunkSize := (numMails + numCPU - 1) / numCPU

	for i := 0; i < numMails; i += chunkSize {
		end := i + chunkSize
		if end > numMails {
			end = numMails
		}

		chunks = append(chunks, paths[i:end])
	}
	return chunks

}
