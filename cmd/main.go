package main

import (
	"cart-api/api"
	config "cart-api/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importação do driver PostgreSQL
)

func main() {
	// Carregar configuração
	cfg := config.LoadConfig()

	// Conectar ao banco de dados
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("não foi possível conectar ao banco de dados:", err)
	}

	// Garantir que a conexão foi estabelecida
	if err = db.Ping(); err != nil {
		log.Fatal("não foi possível verificar a conexão ao banco de dados:", err)
	}
	log.Println("Conexão ao banco de dados estabelecida com sucesso.")

	// Passar a conexão para a função SetupRouter
	router := api.SetupRouter(db)

	// Iniciar o servidor na porta 8080
	log.Println("Iniciando o servidor na porta 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("não foi possível iniciar o servidor:", err)
	}
}


