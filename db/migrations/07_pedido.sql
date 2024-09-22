CREATE TABLE IF NOT EXISTS "pedido" (
  id_pedido SERIAL PRIMARY KEY,        -- Coluna id_pedido como chave primária
  id_cliente INT REFERENCES cliente(id_cliente),  -- Coluna id_cliente referenciando a tabela cliente
  carrinho_id INT REFERENCES carrinho(id),
  status VARCHAR(50) NOT NULL,
  total NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


