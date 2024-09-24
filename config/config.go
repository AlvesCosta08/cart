package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

// LoadConfig carrega a configuração com base no ambiente
func LoadConfig() Config {
	// Verifica o diretório atual
	log.Println("Diretório atual:", os.Getenv("PWD"))

	// Carregar o arquivo .env do diretório pai
	log.Println("Tentando carregar o arquivo .env")
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Carregar as configurações
	config := Config{
		DBUser:     getEnv("POSTGRES_USER"),
		DBPassword: getEnv("POSTGRES_PASSWORD"),
		DBName:     getEnv("POSTGRES_DB"),
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
	}

	log.Printf("Configurações carregadas: %+v", config)
	return config
}

// getEnv obtém o valor da variável de ambiente
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("A variável de ambiente %s é obrigatória, mas não está definida", key)
	}
	return value
}








