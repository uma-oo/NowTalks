package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	database "real-time-forum/backend/database"
	Init "real-time-forum/backend/init"
	"real-time-forum/backend/middleware"
)

// insure that the database is always closed

// close the database

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		panic(err)
	}

	err = db.ReadSQL("./database/db.sql")
	if err != nil {
		fmt.Println("err DB 2", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		db.Database.Close()
		fmt.Println("cleaned!!")
		os.Exit(0)
	}()

	mux := Init.InitSetup(db.Database)
	fmt.Println("Listening on: http://localhost:8080")
	http.ListenAndServe(":8080", middleware.NewRateLimitMiddleWare(mux))
}
