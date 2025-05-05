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
	Id        int    `json:"id,omitempty"`
	PostId    int    `json:"post_id"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at,omitempty"`
	Content   string `json:"content"`
}

type ErrorJson struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
