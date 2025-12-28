package main

import (
	"context"
	"fmt"
	"pgx_text/simple_connection"
)

func main() {

	ctx := context.Background()
	conn, err := simple_connection.CheckConnection(ctx)

	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
