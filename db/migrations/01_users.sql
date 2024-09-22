CREATE TABLE IF NOT EXISTS "user" (
  id_user SERIAL PRIMARY KEY,    -- Coluna id_user como chave primária
  nome VARCHAR(100) NOT NULL,    -- Coluna nome para o nome do usuário
  email VARCHAR(100) UNIQUE NOT NULL,
  senha VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
