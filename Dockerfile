# Usa una imagen base de Go que soporte la versión 1.23
FROM golang:1.23

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto al directorio de trabajo
COPY . .

# Descarga las dependencias del módulo Go
RUN go mod tidy

# Compila la aplicación
RUN go build -o main .

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8081

# Comando para ejecutar la aplicación
CMD ["./main"]