package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
INSERT INTO tasks (title, description, completed,created_at)
VALUES ('Домашка','to do homework Aza',false,'2025-12-29 18:16:16' );`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
