package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type database struct {
	conn *pgx.Conn
	ctx  context.Context
}

func newConn(ctx context.Context, path string) (*pgx.Conn, error) {
	return pgx.Connect(ctx, path)
}

func NewDataBase(ctx context.Context, path string) (*database, error) {
	conn, err := newConn(ctx, path)

	if err != nil {
		return &database{}, err
	}

	return &database{ctx: ctx, conn: conn}, nil
}
