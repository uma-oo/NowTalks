package models

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
	Id        int    `json:"id,omitempty"`
	UserId    int    `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
}

type Comment struct {
	Id        int
	PostId    int
	UserId    int
	CreatedAt string
	Content   string
}
