package main

import (
	"fmt"
	"net/http"

	database "real-time-forum/backend/database"
	Init "real-time-forum/backend/init"
)

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		fmt.Println("err DB", err)
	}

	err = db.ReadSQL("./database/db.sql")
	if err != nil {
		fmt.Println("err DB", err)
	}
	Init.InitSetup(db.Database)

	defer db.Database.Close()

	fmt.Println("Listening on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
