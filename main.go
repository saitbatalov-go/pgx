package main

import (
	"context"
	"fmt"
	"pgx_text/simple_connection"
	"pgx_text/simple_sql"
)

func main() {

	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)

	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.InsertRow(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("Successfully")
}
