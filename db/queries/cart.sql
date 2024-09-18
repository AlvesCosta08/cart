-- db/queries/cart.sql

-- name: CreateCart :one
INSERT INTO carts (user_id) VALUES ($1) RETURNING id;

-- name: AddItemToCart :exec
INSERT INTO cart_items (cart_id, product_id, quantity, price) VALUES ($1, $2, $3, $4);

-- name: ListCartItems :many
SELECT * FROM cart_items WHERE cart_id = $1;

-- name: TotalCartItems :one
SELECT COUNT(*) FROM cart_items WHERE cart_id = $1;

-- name: TotalCartValue :one
SELECT SUM(quantity * price) FROM cart_items WHERE cart_id = $1;

