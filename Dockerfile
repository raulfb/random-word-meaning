# Usa una imagen base de Go para la construcción
FROM golang:1.23 AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Descarga las dependencias del módulo Go
RUN go mod tidy

# Compila la aplicación para Linux en arquitectura ARM
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main .

# Usa una imagen más ligera para el contenedor final
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el binario compilado desde la etapa de construcción
COPY --from=builder /app/main .

# Copia el archivo words.txt al contenedor
COPY --from=builder /app/words.txt .

# Copia la carpeta assets al contenedor
COPY --from=builder /app/assets ./assets

# Copia la carpeta templates al contenedor
COPY --from=builder /app/templates ./templates

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8081

# Comando para ejecutar la aplicación
CMD ["./main"]