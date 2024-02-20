package getenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadEnv() {
	fd, err := os.Open("vars.env")
	if err != nil {
		fmt.Println("Error al leer el archivo de ambientes")
	}

	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Se cargaron las variables de entorno correctamente")
			break
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		} else {
			fmt.Printf("malformed variable line: %s", line)
		}
	}
}
