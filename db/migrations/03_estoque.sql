CREATE TABLE IF NOT EXISTS "estoque" (
  id SERIAL PRIMARY KEY,
  id_produto INT REFERENCES produto(id_produto),  -- Coluna id_produto referenciando a tabela produto
  quantidade_atual INT NOT NULL,          -- Coluna quantidade_atual para a quantidade no estoque
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


