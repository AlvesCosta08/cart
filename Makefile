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
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/01_users.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/04_categoria.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/02_produtos.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/03_estoque.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/08_cliente.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/05_carrinho.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/06_itens_no_carrinho.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/07_pedido.sql	
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/09_documento.sql
	docker-compose exec -T $(DB_CONTAINER_NAME) psql -U postgres -d shopping_cart -f /docker-entrypoint-initdb.d/10_taxas.sql

# Gera o código do SQLC
sqlc:
	sqlc generate

# Rodar todos os comandos necessários
all: up migrate sqlc



