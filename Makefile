# Nome do contêiner PostgreSQL
DB_CONTAINER_NAME = db

# Cria e sobe os containers Docker
up:
	docker-compose up -d

# Destrói os containers Docker
down:
	docker-compose down

# Executa migrações no banco de dados com checagem de tabelas
migrate:	

	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/10_taxas.sql
# Gera o código do SQLC
sqlc:
	sqlc generate

# Rodar todos os comandos necessários
all: up migrate sqlc





