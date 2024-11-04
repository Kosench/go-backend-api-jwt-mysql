package main

import (
	"github.com/go-sql-driver/mysql"
	"go-backend-api-jwt-mysql/cmd/api"
	"go-backend-api-jwt-mysql/config"
	"go-backend-api-jwt-mysql/db"
	"log"
)

func main() {
	storage, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", storage)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
