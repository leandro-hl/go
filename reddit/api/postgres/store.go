package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	api "hl/reddit"
)

const (
	ThreadQuery        = `SELECT * FROM threads WHERE id = $1`
	ThreadsQuery       = `SELECT * FROM threads`
	ThreadInsert       = `INSERT INTO threads VALUES ($1, $2, $3) RETURNING id`
	ThreadUpdate       = `UPDATE threads SET title = $1, description = $2 WHERE id = $3`
	ThreadDelete       = `DELETE FROM threads WHERE id = $1`
	PostQuery          = `SELECT * FROM posts WHERE id = $1`
	PostsByThreadQuery = `
		SELECT
			p.*,
			COUNT(comments.*) AS comments_count,
			t.title AS thread_title
		FROM posts p
		INNER JOIN threads t ON t.id = p.thread_id
		LEFT JOIN comments c ON c.post_id = p.id
		WHERE p.thread_id = $1
		GROUP BY p.id, t.title
		ORDER BY p.votes DESC`
	PostsQuery = `
		SELECT
			p.*,
			COUNT(comments.*) AS comments_count,
			t.title AS thread_title
		FROM posts p
		INNER JOIN threads t ON t.id = p.thread_id
		LEFT JOIN comments c ON c.post_id = p.id
		GROUP BY p.id, t.title
		ORDER BY p.votes DESC`
	PostInsert          = `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING id`
	PostUpdate          = `UPDATE posts SET thread_id = $1, title = $2, content = $3, votes = $4 WHERE id = $5`
	PostDelete          = `DELETE FROM posts WHERE id = $1`
	CommentQuery        = `SELECT * FROM comments WHERE id = $1`
	CommentsByPostQuery = `SELECT * FROM comments WHERE post_id = $1 ORDER BY votes DESC`
	CommentInsert       = `INSERT INTO comments VALUES ($1, $2, $3, $4) RETURNING id`
	CommentUpdate       = `UPDATE comments SET post_id = $1, content = $2, votes = $3 WHERE id = $4`
	CommentDelete       = `DELETE FROM comments WHERE id = $1`
	UserQuery           = `SELECT * FROM users WHERE id = $1`
	UsersQuery          = `SELECT * FROM users`
	UsersByNameQuery    = `SELECT * FROM users WHERE username = $1`
	UserInsert          = `INSERT INTO users VALUES ($1, $2, $3) RETURNING id`
	UserUpdate          = `UPDATE users SET username = $1, password = $2 WHERE id = $3`
	UserDelete          = `DELETE FROM users WHERE id = $1`
)

type Store struct {
	db *sqlx.DB
}

func (s Store) GetPost(id uuid.UUID) (*api.Post, error) {
	panic("implement me")
}

func (s Store) GetPosts() ([]*api.Post, error) {
	panic("implement me")
}

func (s Store) PostsByThread(threadID uuid.UUID) ([]*api.Post, error) {
	panic("implement me")
}

func (s Store) CreatePost(t *api.Post) error {
	panic("implement me")
}

func (s Store) UpdatePost(t *api.Post) error {
	panic("implement me")
}

func (s Store) DeletePost(id uuid.UUID) error {
	panic("implement me")
}

func (s Store) GetThread(id uuid.UUID) (*api.Thread, error) {
	panic("implement me")
}

func (s Store) GetThreads() ([]*api.Thread, error) {
	panic("implement me")
}

func (s Store) CreateThread(t *api.Thread) error {
	panic("implement me")
}

func (s Store) UpdateThread(t *api.Thread) error {
	panic("implement me")
}

func (s Store) DeleteThread(id uuid.UUID) error {
	panic("implement me")
}

func (s Store) GetComment(id uuid.UUID) (*api.Comment, error) {
	panic("implement me")
}

func (s Store) GetCommentsByPost(postID uuid.UUID) ([]*api.Comment, error) {
	panic("implement me")
}

func (s Store) CreateComment(t *api.Comment) error {
	panic("implement me")
}

func (s Store) UpdateComment(t *api.Comment) error {
	panic("implement me")
}

func (s Store) DeleteComment(id uuid.UUID) error {
	panic("implement me")
}

func (s Store) GetUser(id uuid.UUID) (*api.User, error) {
	panic("implement me")
}

func (s Store) GetUserByUsername(username string) (*api.User, error) {
	panic("implement me")
}

func (s Store) CreateUser(u *api.User) error {
	panic("implement me")
}

func (s Store) UpdateUser(u *api.User) error {
	panic("implement me")
}

func (s Store) DeleteUser(id uuid.UUID) error {
	panic("implement me")
}
