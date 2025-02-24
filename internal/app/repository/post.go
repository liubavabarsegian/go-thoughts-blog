package repository

import (
	"context"
	"database/sql"
	"day06/internal/app/models"
	"fmt"
	"log"
	"time"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

func (m *PostRepository) GetByID(ctx context.Context, id int64) (res models.Post, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID = ?`

	fmt.Println(query)
	// list, err := m.fetch(ctx, query, id)
	// if err != nil {
	// 	return models.Post{}, err
	// }

	// if len(list) > 0 {
	// 	res = list[0]
	// } else {
	// 	return res, models.ErrNotFound
	// }

	return
}

func (m *PostRepository) Create(postData models.Post, db *sql.DB) (post models.Post, err error) {
	query := `INSERT INTO posts (title, content, created_at) VALUES ($1, $2, $3) RETURNING id;`

	result, err := db.Exec(query, postData.Title, postData.Content, time.Now())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RESULT", result)
	return models.Post{}, err
}

func (m *PostRepository) GetAllPosts() ([]models.Post, error) {
	query := `SELECT id, title, content, created_at FROM posts;`

	rows, err := m.db.Query(query)
	if rows != nil {
		defer rows.Close()
	}

	posts := []models.Post{}
	// Обрабатываем каждую строку
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RESULT", posts)
	return posts, nil
}
