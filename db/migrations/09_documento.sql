CREATE TABLE IF NOT EXISTS "documento" (
  id SERIAL PRIMARY KEY,
  cliente_id INT REFERENCES cliente(id_cliente),
  tipo_documento VARCHAR(50) NOT NULL,
  numero_documento VARCHAR(50) NOT NULL,
  cpf VARCHAR(14),  -- Adicionando a coluna CPF (opcional ou obrigatório, dependendo da regra)
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

