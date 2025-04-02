package main

import (
	"log"
	"os"
	"project8/telnet/server"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Error: InvalidArgument: Пример верного запуска программы: ./server localhost:8080")
	}

	address := os.Args[1]
	server.Start(address)
}
