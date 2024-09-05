package main

import (
	"log"

	"github.com/Edu4rdoNeves/Client-Server-API/database"
	"github.com/Edu4rdoNeves/Client-Server-API/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Fail to load .env file")
	}

	database.StartDataBaseConnect()
	server.Run()
}
