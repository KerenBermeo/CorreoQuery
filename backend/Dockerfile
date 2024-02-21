FROM golang:1.21.2-alpine

# Instala tini y bash
RUN apk update && apk add tini bash

WORKDIR /app

# Copiar solo los archivos necesarios para instalar dependencias
COPY go.mod .
COPY go.sum .


# Instala dependencias
RUN go mod download

# Copiar el resto de los archivos
COPY . .

# Compila la aplicación
RUN go build -o ./backend/main

# Ejecutar la aplicación
CMD ["./backend/main"]