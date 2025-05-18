package models

import "time"

type User struct {
	Id            int    `json:"id,omitempty"`
	Nickname      string `json:"nickname"`
	Age           int    `json:"age"`
	Gender        string `json:"gender"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	VerifPassword string `json:"password2"`
	CreatedAt     string `json:"created_at,omitempty"`
}


type Session struct {
	Id      int       `json:"id,omitempty"`
	Token   string    `json:"token"`
	UserId  int       `json:"user_id"`
	ExpDate time.Time `json:"expiration_date,omitempty"`
}

type Post struct {
	Id             int    `json:"id,omitempty"`
	UserId         int    `json:"user_id,omitempty"`
	Username       string `json:"user_name,omitempty"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	PostCategories []any  `json:"category_ids"`
	CreatedAt      string `json:"created_at,omitempty"`
	TotalComments  int    `json:"total_comments"`
	TotalLikes     int    `json:"total_likes"`
	TotalDislikes  int    `json:"total_dislikes"`
}

type Category struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Comment struct {
	Id        int    `json:"id,omitempty"`
	PostId    int    `json:"post_id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	Username  string `json:"user_name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Content   string `json:"content"`
	TotalLikes     int    `json:"total_likes"`
	TotalDislikes  int    `json:"total_dislikes"`
}

type PostError struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Categories string  `json:"category_ids"`
}

type CommentError struct {
	Content string `json:"content"`
}

type RegisterError struct {
	Nickname      string `json:"user_name"`
	Age           string `json:"age"`
	Gender        string `json:"gender"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	VerifPassword string `json:"verifpassword"`
}

type Login struct {
	LoginField string `json:"login"`
	Password   string `json:"password"`
}

type UserData struct {
	IsLoggedIn bool   `json:"is_logged_in"`
	Id         int    `json:"id,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
}

// we can make the message interface and then accpet all of them but for now let's work so

type ErrorJson struct {
	Status  int `json:"status"`
	Message any `json:"errors"`
}

//

func NewErrorJson(status int, message any) *ErrorJson {
	return &ErrorJson{
		Status:  status,
		Message: message,
	}
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

func NewSession() *Session {
	return &Session{}
}

func (session *Session) IsExpired() bool {
	return session.ExpDate.Before(time.Now())
}
