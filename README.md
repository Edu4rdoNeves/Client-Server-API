# 💸 Dollar Exchange Rate API - Consulta de Cotação do Dólar

**Bem-vindo à API de Cotação do Dólar!**  
Esta API simples em Go fornece a cotação atual do dólar de maneira rápida e acessível. Você pode consultar a cotação a qualquer momento, e a rota está aberta para consumo sem necessidade de autenticação.

<br>

## 📝 Descrição

A **Dollar Exchange Rate API** permite que qualquer usuário consulte a cotação atual do dólar em relação à moeda local. Esta API foi construída usando Golang e segue práticas RESTful, oferecendo uma interface limpa e fácil de usar.

Além disso, a cotação mais recente é **salva automaticamente em um arquivo `.txt`** com o valor e a data, permitindo que você faça manipulações ou análises futuras da cotação de maneira fácil.

<br>

## 🚀 Como Funciona?

A API possui apenas uma rota:

- **GET** `/api/v1/quote`

Ao fazer uma solicitação para essa rota, você receberá como resposta a cotação atual do dólar. Além disso, a cotação é salva em um arquivo chamado `latest_quote.txt` no diretório do projeto, contendo o valor e o timestamp.

### Exemplo de Requisição:
```bash
GET http://localhost:8080/api/v1/quote
```
### Exemplo de Resposta:

```json
{
  "bid": 5.20,
  "created_at": "2024-09-02T12:45:00Z"
}
```
### Exemplo de Arquivo `latest_rate.txt`:

```txt
Cotação do Dólar: 5.20
Data: 2024-09-02T12:45:00Z
```

<br>

## 🛠️ Como Executar o Projeto Localmente

Se você deseja rodar a API localmente, siga os passos abaixo:

### Pré-requisitos:

- Go 1.18 ou superior
- Docker (opcional, caso queira usar um ambiente containerizado)

### Passos:

1. Clone o repositório:

```bash
git clone https://github.com/Edu4rdoNeves/Client-Server-API.git
```
2. Acesse a pasta do projeto:

```bash
cd seu-repositorio
```

3. Instale as dependências:
```bash
go mod tidy
```

4. Execute o servidor:
```bash
go run main.go
```

5. Acesse a API:
```bash
http://localhost:8080/api/v1/quote
```

<br>

## 📦 Executando com Docker

Você também pode rodar a aplicação usando Docker:

### 1. Crie a imagem do Docker:

```bash
docker build -t dollar-quote-api .
```

### 2. Execute o container:

```bash
docker run -p 8080:8080 dollar-quote-api
```

### 3. Acesse a API no navegador ou via Postman:

```bash
http://localhost:8080/api/v1/quote.
```
<br>

## 📚 Documentação Completa

### Rota Disponível

- **GET** `/api/v1/quote`:
  - **Descrição**: Retorna a cotação atual do dólar e salva a cotação mais recente no arquivo `latest_quote.txt`.
  - **Resposta**:
    - **currency**: Sempre retorna "USD" (dólar americano).
    - **rate**: Cotação atual do dólar.
    - **timestamp**: Data e hora da consulta no formato ISO 8601.

### Código de Status

- **200 OK**: A requisição foi bem-sucedida.
- **500 Internal Server Error**: Ocorreu um erro no servidor ao tentar obter a cotação.

<br>

## 🤝 Contribuição

Contribuições são bem-vindas! Se você quiser melhorar ou adicionar novos recursos a esta API, sinta-se à vontade para abrir um PR ou enviar uma issue.

<br>

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](./LICENSE) para mais detalhes.

<br>

## 🛠️ Tecnologias Utilizadas

- [Golang](https://golang.org/) - Linguagem de programação usada para construir a API.
- [Gin Gonic](https://gin-gonic.com/) - Framework web leve para Go.
- [Gorm](https://gorm.io/) - ORM (Object-Relational Mapping) para Go.
- [Docker](https://www.docker.com/) - Para containerização da aplicação.

<br>
<br>
<br>

Feito com ❤️ por [Eduardo Neves](https://github.com/Edu4rdoNeves)













