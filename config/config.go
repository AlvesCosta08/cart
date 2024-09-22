package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // Define o padrão como "local"
	}

	// Carregar o arquivo .env correspondente
	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Erro ao carregar o arquivo %s: %v", envFile, err)
	}

	// Carregar as configurações
	config := Config{
		DBUser:     getEnv("POSTGRES_USER"),
		DBPassword: getEnv("POSTGRES_PASSWORD"),
		DBName:     getEnv("POSTGRES_DB"),
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
	}

	log.Printf("Configurações carregadas para o ambiente: %s", env)
	return config
}

// getEnv obtém o valor da variável de ambiente
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}





