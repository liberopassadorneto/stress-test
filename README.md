# Stress-Test

## Visão Geral

**Stress-Test** é uma ferramenta de linha de comando (CLI) desenvolvida em Go para realizar testes de carga em serviços web. Permite que os usuários especifiquem a URL do serviço alvo, o número total de requisições e o nível de concorrência (número de chamadas simultâneas). Após a execução dos testes, a ferramenta gera um relatório detalhado com métricas essenciais para avaliar o desempenho do serviço.

## Funcionalidades

- **Configuração Personalizada**: Especifique a URL do serviço, o número total de requisições e o nível de concorrência.
- **Execução Concorrente**: Realiza requisições HTTP de forma simultânea, simulando múltiplos usuários acessando o serviço ao mesmo tempo.
- **Relatório Detalhado**: Gera um relatório que inclui tempo total de execução, número de requisições realizadas, sucesso das requisições (status 200) e distribuição de outros códigos de status HTTP.
- **Containerização com Docker**: Facilita a execução da ferramenta em diferentes ambientes sem necessidade de instalar dependências localmente.

## Estrutura do Projeto

```
stress-test/
├── cmd/
│   └── main.go
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Instalação

### Pré-requisitos

- [Go](https://golang.org/dl/) 1.20 ou superior
- [Docker](https://www.docker.com/get-started) (opcional, para execução via container)

### 1. Clonar o Repositório

```bash
git clone <https://github.com/liberopassadorneto/stress-test>
cd stress-test
```

### 2. Compilar e Executar Localmente

#### a. Configurar o Módulo Go

Caso ainda não tenha inicializado o módulo Go, faça isso:

```bash
go mod init stress-test
go mod tidy
```

#### b. Compilar a Aplicação

```bash
go build -o stress-test ./cmd
```

#### c. Executar a Aplicação

```bash
./stress-test --url=http://google.com --requests=1000 --concurrency=10
```

### 3. Executar com Docker

#### a. Construir a Imagem Docker

Certifique-se de estar no diretório raiz do projeto (`stress-test/`).

```bash
docker build -t stress-test .
```

#### b. Executar o Container

```bash
docker run --rm stress-test --url=http://google.com --requests=1000 --concurrency=10
```

## Uso

A ferramenta aceita os seguintes parâmetros via linha de comando:

- `--url` (obrigatório): URL do serviço a ser testado.
- `--requests` (obrigatório): Número total de requisições a serem enviadas.
- `--concurrency` (opcional): Número de chamadas simultâneas (padrão: 1).

### Sintaxe

```bash
stress-test --url=<URL_DO_SERVIÇO> --requests=<NÚMERO_TOTAL_DE_REQUESTS> --concurrency=<NÍVEL_DE_CONCORRÊNCIA>
```

### Exemplos

#### Exemplo 1: Testar o Google com 1000 requisições e 10 chamadas simultâneas

```bash
./stress-test --url=http://google.com --requests=1000 --concurrency=10
```

**Saída Esperada:**

```
===== Relatório de Teste de Carga =====
URL Testada: http://google.com
Total de Requests: 1000
Concorrência: 10
Tempo Total: 5.432s
Requests com Status 200: 1000
Distribuição de Status Codes:
  200: 1000
```

#### Exemplo 2: Testar um Serviço Local com 500 requisições e 50 chamadas simultâneas

```bash
./stress-test --url=http://localhost:8080/api --requests=500 --concurrency=50
```

**Saída Esperada:**

```
===== Relatório de Teste de Carga =====
URL Testada: http://localhost:8080/api
Total de Requests: 500
Concorrência: 50
Tempo Total: 3.210s
Requests com Status 200: 495
Distribuição de Status Codes:
  200: 495
  404: 5
```

## Explicação dos Resultados

- **Tempo Total**: Tempo total gasto na execução dos testes de carga.
- **Requests com Status 200**: Número de requisições que retornaram com sucesso (HTTP 200).
- **Distribuição de Status Codes**: Quantidade de requisições por código de status HTTP. Erros de conexão são contabilizados como `0`.