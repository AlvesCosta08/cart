CREATE TABLE IF NOT EXISTS "cliente" (
  id_cliente SERIAL PRIMARY KEY,      -- Coluna id_cliente como chave primária
  nome_fantasia VARCHAR(100),         -- Coluna nome_fantasia (opcional)
  name VARCHAR(100) NOT NULL,         -- Nome do cliente
  email VARCHAR(100) UNIQUE NOT NULL, -- E-mail único e obrigatório
  telefone VARCHAR(15),
  endereco TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


