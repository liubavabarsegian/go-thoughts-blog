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

func (repo *PostRepository) GetByID(ctx context.Context, id int64) (res models.Post, err error) {
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

func (repo *PostRepository) CreatePost(postData models.Post) (post models.Post, err error) {
	query := `INSERT INTO posts (title, content, created_at) VALUES ($1, $2, $3) RETURNING id;`

	err = repo.db.QueryRow(query, postData.Title, postData.Content, time.Now()).Scan(&postData.ID)
	if err != nil {
		return models.Post{}, err
	}

	return postData, nil
}

func (repo *PostRepository) GetAllPosts() ([]models.Post, error) {
	query := `SELECT id, title, content, created_at FROM posts;`

	rows, err := repo.db.Query(query)
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
