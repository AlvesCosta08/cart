# Shopping Cart API

Este projeto implementa uma API para gerenciamento de um carrinho de compras utilizando Golang, PostgreSQL, SQLC e Docker.

## Estrutura do Projeto

```bash
shopping-cart-api/
├── api
│   └── handlers
│       └── product.go     # Handlers do Gin
├── db
│   ├── migrations
│   │   └── 001_create_products_table.sql   # Migrações do banco de dados
│   └── queries
│       └── products.sql   # Consultas SQL utilizadas pelo SQLC
├── internal
│   └── db
│       └── models.go      # Código gerado pelo SQLC
├── docker-compose.yml      # Arquivo para subir PostgreSQL via Docker
├── sqlc.yaml               # Configuração do SQLC
├── Makefile                # Arquivo Makefile para automação de tarefas
├── go.mod                  # Arquivo de dependências do Go
├── go.sum                  # Hash das dependências
└── main.go                 # Entrada principal da aplicação
