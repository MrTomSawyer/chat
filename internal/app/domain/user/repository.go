package user

import (
	"context"
	"fmt"

	"github.com/MrTomSawyer/chat/internal/app/model"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db    *pgx.Conn
	table string
}

func NewUserRepository(db *pgx.Conn, table string) *Repository {
	return &Repository{
		db:    db,
		table: table,
	}
}

func (r *Repository) Create(ctx context.Context, user *model.User) error {
	q := fmt.Sprintf("INSERT INTO %s (id, name, password_hash) VALUES ($1, $2, $3)", r.table)
	_, err := r.db.Exec(ctx, q, user.ID, user.Name, user.PasswordHash)
	return err
}
