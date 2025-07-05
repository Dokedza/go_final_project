package main

import (
	"fmt"
	database "go1f/pkg/db"
	"go1f/pkg/server"
)

func main() {

	dbFile := "scheduler.db"

	err := database.Init(dbFile)
	if err != nil {
		fmt.Println(err)
	}
	err = server.Run()
	if err != nil {
		panic(err)
	}

}
