FROM golang:1.24-alpine

WORKDIR /app

# Instalar dependências do sistema
RUN apk add --no-cache git

# Copiar mod files
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN go build -o main ./cmd/api

# Expor porta
EXPOSE 8080

# Comando de execução
CMD ["./main"]