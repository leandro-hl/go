package api

import (
	"github.com/google/uuid"
)

type Thread struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type Post struct {
	ID            uuid.UUID `json:"id"`
	ThreadID      uuid.UUID `json:"thread_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Votes         int       `json:"votes"`
	CommentsCount int       `json:"comments_count"` //not in posts table
	ThreadTitle   string    `json:"thread_title"`   //not in posts table
}

type Comment struct {
	ID      uuid.UUID `json:"id"`
	PostID  uuid.UUID `json:"post_id"`
	Content string    `json:"content"`
	Votes   int       `json:"votes"`
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type Store interface {
	GetPost(id uuid.UUID) (*Post, error)
	GetPosts() ([]*Post, error)
	PostsByThread(threadID uuid.UUID) ([]*Post, error)
	CreatePost(t *Post) error
	UpdatePost(t *Post) error
	DeletePost(id uuid.UUID) error
	GetThread(id uuid.UUID) (*Thread, error)
	GetThreads() ([]*Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
	GetComment(id uuid.UUID) (*Comment, error)
	GetCommentsByPost(postID uuid.UUID) ([]*Comment, error)
	CreateComment(t *Comment) error
	UpdateComment(t *Comment) error
	DeleteComment(id uuid.UUID) error
	GetUser(id uuid.UUID) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(id uuid.UUID) error
}
