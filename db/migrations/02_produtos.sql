CREATE TABLE IF NOT EXISTS "produto" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price NUMERIC(10, 2) NOT NULL,
  categoria_id INT REFERENCES categoria(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

