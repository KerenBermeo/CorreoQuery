package data

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/mail"
// 	"os"
// 	"strings"

// 	"github.com/KerenBermeo/CorreoQuery/model"
// )

// func ParseEmailFromPath(contador int, paths []string) ([]model.Email, error) {
// 	var emails []model.Email
// 	archivoNum := 0
// 	fmt.Println("PARSEANDO LOTE NUMERO ", contador, " CON ", len(paths), " ARCHIVOS")
// 	fmt.Println("---------------(#LOTE/#ARCHIVO)---------------------")
// 	for _, filePath := range paths {
// 		archivoNum++
// 		fmt.Println("---------------(  ", contador, "  |  ", archivoNum, " )------------------ ")
// 		fd, err := os.Open(filePath)
// 		if err != nil {
// 			log.Printf("Error al abrir el archivo %s: %v", filePath, err)
// 			continue // Continuar con el próximo archivo
// 		}

// 		// Cerrar el archivo después de leerlo
// 		defer fd.Close()

// 		// Leer y parsear el mensaje de correo electrónico
// 		m, err := mail.ReadMessage(fd)
// 		if err != nil {
// 			log.Printf("Error al analizar el archivo %s: %v", filePath, err)
// 			continue // Continuar con el próximo archivo
// 		}

// 		header := m.Header
// 		body, err := io.ReadAll(m.Body)
// 		if err != nil {
// 			log.Printf("Error al leer el cuerpo del correo electrónico %s: %v", filePath, err)
// 			continue // Continuar con el próximo archivo
// 		}

// 		recipients := strings.Split(header.Get("To"), ", ")

// 		// Crear el correo electrónico solo si todas las operaciones fueron exitosas
// 		email := model.Email{
// 			MessageId: header.Get("Message-ID"),
// 			Date:      header.Get("Date"),
// 			From:      header.Get("From"),
// 			To:        recipients,
// 			Subject:   header.Get("Subject"),
// 			Content:   string(body),
// 		}
// 		emails = append(emails, email)
// 	}

// 	//fmt.Println(emails)

// 	return emails, nil
// }

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KerenBermeo/CorreoQuery/model"
)

func ParseEmailFromPath(contador int, paths []string) ([]model.Email, error) {

	var email []model.Email
	archivoNum := 0
	// fmt.Println("PARSEANDO LOTE NUMERO ", contador, " CON ", len(paths), " ARCHIVOS")
	// fmt.Println("---------------(#LOTE/#ARCHIVO)---------------------")
	for _, path := range paths {
		archivoNum++
		// fmt.Println("---------------(  ", contador, "  |  ", archivoNum, " )------------------ ")
		fd, err := os.Open(path)
		if err != nil {
			return []model.Email{}, fmt.Errorf("failed to open file: %v", err)
		}
		defer fd.Close()

		scanner := bufio.NewScanner(fd)

		var headerTextBuffer strings.Builder
		var bodyTextBuffer strings.Builder

		// Leer el encabezado del correo electrónico
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break // La línea vacía marca el final del encabezado
			}
			headerTextBuffer.WriteString(line + "\n")
		}

		// Verificar errores durante la lectura del encabezado
		if err := scanner.Err(); err != nil {
			if err != bufio.ErrTooLong {
				return []model.Email{}, fmt.Errorf("error while reading email header: %v in folder %v", err, path)
			}
			continue // Ignorar el archivo que no se puede leer completamente y continuar con el siguiente
		}

		// Leer el cuerpo del correo electrónico
		for scanner.Scan() {
			bodyTextBuffer.WriteString(scanner.Text() + "\n")
		}

		// Verificar errores durante la lectura del cuerpo
		if err := scanner.Err(); err != nil {
			if err != bufio.ErrTooLong {
				return []model.Email{}, fmt.Errorf("error while reading email body: %v in folder %v", err, path)
			}
			continue // Ignorar el archivo que no se puede leer completamente y continuar con el siguiente
		}

		// Parsear el correo electrónico

		parsedEmail, err := ParseEmail(headerTextBuffer.String(), bodyTextBuffer.String())

		if err != nil {
			return []model.Email{}, fmt.Errorf("error parsing email in folder %v: %v", path, err)
		}

		email = append(email, parsedEmail)

	}
	return email, nil
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

	// Verificar campos requeridos
	// requiredFields := []string{"Message-ID", "From", "To", "Subject", "Date"}
	// for _, field := range requiredFields {
	// 	if _, ok := header[field]; !ok {
	// 		return model.Email{}, fmt.Errorf("missing required field in header: %s", field)
	// 	}
	// }

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
