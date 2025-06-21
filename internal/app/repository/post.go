package repository

import (
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

func (repo *PostRepository) GetPost(id uint) (post models.Post, err error) {
	query := `SELECT id, title, content, created_at FROM posts WHERE ID = $1`

	row := repo.db.QueryRow(query, id)
	err = row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
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

func (repo *PostRepository) GetAllPosts(page uint) ([]models.Post, error) {
	query := `SELECT id, title, content, created_at FROM posts LIMIT $1 OFFSET $2;`

	offset := (page - 1) * 3
	rows, err := repo.db.Query(query, 3, offset)
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
		return nil, fmt.Errorf("GetAllPosts Error: ", err)
	}
	fmt.Println("RESULT", posts)
	return posts, nil
}
