# Usar a imagem oficial do Go para construir a aplicação
FROM golang:1.22.1-alpine AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código da aplicação para o diretório de trabalho
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fiap-rocks ./cmd/kitchencontrol/main.go

#Etapa 2: Imagem final mínima para execução
#FROM alpine:latest
FROM golang:1.22.1-alpine

#Adicionar bibliotecas necessárias para o MySQL
#RUN apk add --no-cache libc6-compat mariadb-connector-c

#Definir o diretório de trabalho
WORKDIR /root/

#Copiar o binário da aplicação da etapa de build
COPY --from=builder /app/fiap-rocks .

# Expor a porta em que a aplicação vai rodar
EXPOSE 8083

# Comando para rodar a aplicação
CMD ["./fiap-rocks"]
