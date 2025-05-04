package models

import "time"

type User struct {
	Id        int
	Nickname  string
	Age       int
	Gender    string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Session struct {
	Id    int
	Token string
}

type Post struct {
	Id        int `json:"id"`
	UserId    int
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`

}

type Comment struct {
	Id      int
	PostId int
	UserId int
	CreatedAt time.Time
	Content string
}
