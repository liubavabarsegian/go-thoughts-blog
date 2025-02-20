package post

import (
	"database/sql"
	models "day06/internal/app/models"
	"fmt"
	"log"
	"time"
)

func Create(postData models.Post, db *sql.DB) (post models.Post, err error) {
	query := `INSERT INTO posts (title, content, created_at) VALUES ($1, $2, $3) RETURNING id;`

	result, err := db.Exec(query, postData.Title, postData.Content, time.Now())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RESULT", result)
	return models.Post{}, err
}

func GetAll(db *sql.DB) []models.Post {
	query := `SELECT id, title, content, created_at FROM posts;`

	rows, err := db.Query(query)
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
	return posts
}
