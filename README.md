
# API de CRUD de Produtos em Go

Esta é uma API simples em Go para realizar operações CRUD (Create, Read, Update, Delete) em produtos, conectando-se a um banco de dados PostgreSQL.

## Como Usar

### 1. Clonar o Repositório

Primeiro, clone o repositório para sua máquina local:

```bash
git clone https://github.com/seu-usuario/goapi.git
cd goapi
```

### 2. Baixar e Rodar o Banco de Dados PostgreSQL

Você pode configurar o banco de dados utilizando Docker. Siga os passos abaixo para rodá-lo:

#### 2.1. Instalar o Docker

Se você ainda não tem o Docker instalado, siga o guia oficial para instalar em seu sistema:  
[Guia de instalação do Docker](https://docs.docker.com/get-docker/)

#### 2.2. Subir o Banco de Dados com Docker

Use o Docker para rodar o banco de dados PostgreSQL com o seguinte comando:

```bash
docker-compose up -d go_db
```

Isso criará um container com PostgreSQL rodando na porta `5432`, utilizando as variáveis de ambiente padrões do arquivo `docker-compose.yml` (com senha genérica, altere se necessário).

### 3. Configurar o Banco de Dados (Opcional)

Se você preferir usar um banco de dados local em vez do Docker, ou configurar seu próprio PostgreSQL, siga as instruções abaixo:

#### 3.1. Baixar e Instalar o PostgreSQL

Baixe e instale o PostgreSQL em sua máquina a partir do site oficial:  
[Download do PostgreSQL](https://www.postgresql.org/download/)

#### 3.2. Criar o Banco de Dados

Crie um banco de dados no PostgreSQL com as seguintes credenciais:

- **Usuário**: `seu_usuario`
- **Senha**: `sua_senha`
- **Banco de Dados**: `seu_banco_de_dados`

#### 3.3. Atualizar as Configurações do Banco de Dados

Atualize as configurações do banco de dados no arquivo `db.go` com suas credenciais ou mantenha as configurações padrões. Aqui estão as variáveis para modificar:

```go
const (
    host     = "localhost"
    port     = "5432"
    user     = "postgres"
    password = "1234"
    dbname   = "postgres"
)
```

### 4. Rodar a API

Depois de configurar o banco de dados, rode a API com o comando abaixo:

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

### 5. Testar os Endpoints

A API oferece os seguintes endpoints:

- **GET /products**: Retorna todos os produtos.
- **POST /product**: Cria um novo produto.
- **GET /product/:productId**: Retorna um produto pelo ID.
- **DELETE /product/:productId**: Deleta um produto pelo ID.
