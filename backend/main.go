package main

import (
	"fmt"

	database "real-time-forum/backend/database"
)

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		fmt.Println("err", err)
	}

	err = db.ReadSQL("./database/db.sql")
	if err != nil {
		fmt.Println("err", err)
	}

	// setup layers
	// repo := repositories.NewAppRepository(db.Database)
	// service := service.NewPostService(repo)
	// handler := handler.NewPostService(service)

	defer db.Database.Close()

	_, err = db.Database.Exec(`INSERT INTO categories (category) values ("IT");`)
	fmt.Println("res", err)
}
