// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cliente.sql

package db

import (
	"context"
	"database/sql"
)

const createClient = `-- name: CreateClient :exec
INSERT INTO "cliente" (nome_fantasia, name, email, telefone, endereco) VALUES ($1, $2, $3, $4, $5) RETURNING id_cliente
`

type CreateClientParams struct {
	NomeFantasia sql.NullString
	Name         string
	Email        string
	Telefone     sql.NullString
	Endereco     sql.NullString
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) error {
	_, err := q.db.ExecContext(ctx, createClient,
		arg.NomeFantasia,
		arg.Name,
		arg.Email,
		arg.Telefone,
		arg.Endereco,
	)
	return err
}

const deleteClient = `-- name: DeleteClient :exec
DELETE FROM "cliente" WHERE id_cliente = $1
`

func (q *Queries) DeleteClient(ctx context.Context, idCliente int32) error {
	_, err := q.db.ExecContext(ctx, deleteClient, idCliente)
	return err
}

const getClientByID = `-- name: GetClientByID :one
SELECT id_cliente, nome_fantasia, name, email, telefone, endereco, created_at FROM "cliente" WHERE id_cliente = $1
`

func (q *Queries) GetClientByID(ctx context.Context, idCliente int32) (Cliente, error) {
	row := q.db.QueryRowContext(ctx, getClientByID, idCliente)
	var i Cliente
	err := row.Scan(
		&i.IDCliente,
		&i.NomeFantasia,
		&i.Name,
		&i.Email,
		&i.Telefone,
		&i.Endereco,
		&i.CreatedAt,
	)
	return i, err
}

const updateClient = `-- name: UpdateClient :exec
UPDATE "cliente" SET nome_fantasia = $1, name = $2, email = $3, telefone = $4, endereco = $5 WHERE id_cliente = $6
`

type UpdateClientParams struct {
	NomeFantasia sql.NullString
	Name         string
	Email        string
	Telefone     sql.NullString
	Endereco     sql.NullString
	IDCliente    int32
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) error {
	_, err := q.db.ExecContext(ctx, updateClient,
		arg.NomeFantasia,
		arg.Name,
		arg.Email,
		arg.Telefone,
		arg.Endereco,
		arg.IDCliente,
	)
	return err
}
