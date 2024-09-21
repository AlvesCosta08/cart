CREATE TABLE IF NOT EXISTS "documento" (
  id SERIAL PRIMARY KEY,
  cliente_id INT REFERENCES cliente(id),
  tipo_documento VARCHAR(50) NOT NULL,
  numero_documento VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
