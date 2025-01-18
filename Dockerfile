# Etapa de Build
FROM golang:1.20-alpine AS builder

# Diretório de trabalho dentro do container
WORKDIR /app

# Copiando o go.mod e go.sum para cache de dependências
COPY go.mod go.sum ./

# Baixando as dependências
RUN go mod download

# Copiando o código fonte
COPY . .

# Compilando a aplicação
RUN go build -o stress-test ./cmd

# Etapa Final
FROM alpine:latest

# Diretório de trabalho no container final
WORKDIR /root/

# Copiando o binário compilado do estágio de build
COPY --from=builder /app/stress-test .

# Definindo permissão de execução
RUN chmod +x stress-test

# Comando padrão
ENTRYPOINT ["./stress-test"]
