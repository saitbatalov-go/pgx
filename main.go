package main

import (
	"fmt"
	"pgx_text/http_server"
)

func main() {
	fmt.Println("Запуск HTTP Server круто!")

	err := http_server.StartHTTPServer()
	if err != nil {
		fmt.Println("Произошла ошибка во время запуска сервера =(", err)
	} else {
		fmt.Println("Сервер запустился успешгнл!")
	}
}
