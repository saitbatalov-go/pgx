package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(
	ctx context.Context,
	conn *pgx.Conn,
	// title string,
	// description string,
	// completed bool,
	// createdAt time.Time,
) error {
	sqlQuery := `
UPDATE tasks
SET completed=TRUE
WHERE id=1;
`

	_, err := conn.Exec(ctx, sqlQuery) //title, description, completed, createdAt

	return err
}
