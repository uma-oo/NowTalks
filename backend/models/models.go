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



type PostError struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}


type CommentError struct {
	Content string `json:"content"`
}

// we can make the message interface and then accpet all of them but for now let's work so

type ErrorJson struct {
	Status  int `json:"status"`
	Message any `json:"errors"`
}

//

func NewErrorJson() *ErrorJson {
	return &ErrorJson{}
}

func NewPost() *Post {
	return &Post{}
}

func NewPostErr() *PostError {
	return &PostError{}
}

func NewComment() *Comment {
	return &Comment{}
}

func NewCommentErr() *CommentError {
	return &CommentError{}
}

func NewUser() *User {
	return &User{}
}
