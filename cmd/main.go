package main

import (
	"cart-api/api"
	"cart-api/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	loadEnv()

	// Carregar a configuração
	cfg := config.LoadConfig()

	// Iniciar o servidor da API
	startServer(cfg)
}

// loadEnv carrega as variáveis do arquivo .env
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}

// startServer inicializa o servidor com a configuração fornecida
func startServer(cfg config.Config) {
	router := api.SetupRouter()

	// Inicia o servidor e trata erros
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	log.Printf("Servidor rodando na porta :8080 com o banco: %s", cfg.DBName)
}
