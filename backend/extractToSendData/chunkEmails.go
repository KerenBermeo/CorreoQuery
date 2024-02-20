package data

import "runtime"

// ChunkEmails divide una lista de rutas de correos electrónicos en fragmentos más pequeños.
func ChunkEmails(mailsPaths []string) [][]string {
	// Se declara una variable slices de slices para almacenar los fragmentos de las rutas de correos electrónicos.
	var chunks [][]string

	// Se determina el número de núcleos lógicos disponibles en el sistema.
	numCPU := runtime.NumCPU()

	// Se obtiene la cantidad total de rutas de correos electrónicos en la lista.
	numMails := len(mailsPaths)

	// Se calcula el tamaño de cada fragmento, considerando el número de núcleos lógicos.
	chunkSize := (numMails + numCPU - 1) / numCPU

	// Se itera a través de las rutas de correos electrónicos, creando fragmentos.
	for i := 0; i < numMails; i += chunkSize {
		// Se calcula el índice de finalización del fragmento actual.
		end := i + chunkSize

		// Si el índice de finalización supera la cantidad total de rutas de correos electrónicos,
		// se ajusta para que coincida con la cantidad total.
		if end > numMails {
			end = numMails
		}

		// Se agrega el sub-slice actual al slice de fragmentos.
		chunks = append(chunks, mailsPaths[i:end])
	}

	// Se devuelve el slice de fragmentos resultante.
	return chunks
}
