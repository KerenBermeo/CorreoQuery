package data

import (
	"io"
	"log"
	"net/mail"
	"os"
	"strings"

	"github.com/KerenBermeo/CorreoQuery/model"
)

// parseEmailFromPath lee un archivo de correo electrónico desde la ruta especificada, lo parsea y devuelve una estructura de modelo Email.
// La función toma una ruta de archivo como entrada y devuelve una estructura de modelo Email y un posible error.
func ParseEmailFromPath(path string) (model.Email, error) {
	// Abre el archivo en la ruta especificada.
	fd, err := os.Open(path)

	if err != nil {
		log.Print("Error leyendo el archivo con la ruta", path, err)
		return model.Email{}, err
	}
	defer fd.Close() // Cierra el archivo después de su uso para evitar posibles fugas de recursos.

	// Lee y parsea el mensaje de correo electrónico utilizando el paquete "mail".
	m, err := mail.ReadMessage(fd)

	if err != nil {
		log.Print("Error al analizar el archivo con la ruta", path, err)
		return model.Email{}, err
	}
	header := m.Header // Obtiene el encabezado del mensaje.

	// Lee el cuerpo del mensaje de correo electrónico.
	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Print("Error leyendo el cuerpo del correo electrónico con la ruta", path, err)
		return model.Email{}, err
	}

	// Obtiene los destinatarios del encabezado y los divide en una lista.
	recipients := strings.Split(header.Get("To"), ", ")

	// Crea y devuelve una estructura de modelo Email con la información extraída del correo electrónico parseado.
	mail := model.Email{
		MessageId: header.Get("Message-ID"),
		Date:      header.Get("Date"),
		From:      header.Get("From"),
		To:        recipients,
		Subject:   header.Get("Subject"),
		Content:   string(body),
	}

	return mail, nil
}
