CREATE TABLE IF NOT EXISTS "produto" (
  id_produto SERIAL PRIMARY KEY,    -- Coluna id_produto como chave primária
  name VARCHAR(100) NOT NULL,
  price NUMERIC(10, 2) NOT NULL,
  referencia VARCHAR(100),          -- Coluna referencia para o código do produto
  categoria_id INT REFERENCES categoria(id_categoria),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


