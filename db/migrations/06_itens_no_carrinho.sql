CREATE TABLE IF NOT EXISTS "itens_no_carrinho" (
  id SERIAL PRIMARY KEY,
  cart_id INT REFERENCES carrinho(id),  -- Coluna cart_id referenciando a tabela carrinho
  produto_id INT REFERENCES produto(id_produto),
  quantidade INT NOT NULL,
  preco_unitario NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


