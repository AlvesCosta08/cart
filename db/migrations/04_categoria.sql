CREATE TABLE IF NOT EXISTS "categoria" (
  id_categoria SERIAL PRIMARY KEY,  -- Coluna id_categoria como chave primária
  nome VARCHAR(100) NOT NULL,       -- Coluna nome para o nome da categoria
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


