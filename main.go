package main

import (
	"fmt"

	"github.com/rudrprasad05/go-sql/connect"
)

func main() {
	config := connect.Config{
		Username: "root",
		Password: "",
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "golangtest",
	}
	db, err := connect.InitDB(&config)
	if err != nil{
		fmt.Println("Couldnt connect to db: ", err)
		return
	}
	pingE := db.Ping()
	if pingE != nil{
		fmt.Println("failed to ping", pingE)
		return
	}
}