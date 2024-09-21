CREATE TABLE IF NOT EXISTS "itens_no_carrinho" (
  id SERIAL PRIMARY KEY,
  carrinho_id INT REFERENCES carrinho(id),
  produto_id INT REFERENCES produto(id),
  quantidade INT NOT NULL,
  preco_unitario NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

