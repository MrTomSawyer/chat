package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PG struct {
	Conn *pgx.Conn
}

func NewPG() *PG {
	return &PG{}
}

func (pg *PG) Connect(ctx context.Context, dsn string) error {
	var err error

	pg.Conn, err = pgx.Connect(ctx, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PG) Disconnect(ctx context.Context) error {
	return pg.Conn.Close(ctx)
}
