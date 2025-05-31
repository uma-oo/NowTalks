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
	Id             int       `json:"id,omitempty"`
	UserId         int       `json:"user_id,omitempty"`
	Username       string    `json:"user_name,omitempty"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	PostCategories []any     `json:"categories"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	TotalComments  int       `json:"total_comments"`
	TotalLikes     int       `json:"total_likes"`
	TotalDislikes  int       `json:"total_dislikes"`
}

type Category struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Comment struct {
	Id            int       `json:"id,omitempty"`
	PostId        int       `json:"post_id,omitempty"`
	UserId        int       `json:"user_id,omitempty"`
	Username      string    `json:"user_name,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Content       string    `json:"content"`
	TotalLikes    int       `json:"total_likes"`
	TotalDislikes int       `json:"total_dislikes"`
}

type Reaction struct {
	Id           int    `json:"id,omitempty"`
	EntityTypeId int    `json:"entity_type_id,omitempty"`
	EntityType   string `json:"entity_type,omitempty"`
	EntityId     int    `json:"entity_id"`
	Reaction     int    `json:"reaction,omitempty"`
	UserId       int    `json:"user_id"`
}

type PostError struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Categories string `json:"categories"`
}

type CommentError struct {
	Content string `json:"content"`
}

type RegisterError struct {
	Nickname      string `json:"nickname"`
	Age           string `json:"age"`
	Gender        string `json:"gender"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	VerifPassword string `json:"password2"`
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

//  Message stuff ;)

type Message struct {
	SenderID         int       `json:"sender_id,omitempty"`
	SenderUsername   string    `json:"sender_username,omitempty"`
	ReceiverID       int       `json:"receiver_id"`
	ReceiverUsername string    `json:"receiver_username"`
	Message          string    `json:"message"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	Status           string    `json:"status,omitempty"`
}

type MessageErr struct {
	Message    string `json:"message"`
	ReceiverID any    `json:"receiver_id"`
}

type ReactionErr struct {
	EntityId string `json:"entity_id"`
}




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

func NewMessageErr() *MessageErr {
	return &MessageErr{}
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
