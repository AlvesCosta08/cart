CREATE TABLE IF NOT EXISTS "taxas" (
  id SERIAL PRIMARY KEY,
  cart_id INT REFERENCES carrinho(id),  -- Coluna cart_id referenciando a tabela carrinho
  descricao VARCHAR(100) NOT NULL,
  value NUMERIC(10, 2) NOT NULL,        -- Coluna value para o valor da taxa
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


