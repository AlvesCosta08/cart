-- db/migrations/002_create_cart_table.sql
CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INT REFERENCES carts(id) ON DELETE CASCADE,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);