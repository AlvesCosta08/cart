CREATE TABLE IF NOT EXISTS "carrinho" (
  id SERIAL PRIMARY KEY,
  cliente_id INT REFERENCES cliente(id_cliente),
  user_id INT REFERENCES "user"(id_user),  
  total NUMERIC(10, 2) DEFAULT 0,
  status VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


