package posts

import (
	"database/sql"
	postEntities "day06/internal/app/entities/post"
	"fmt"
	"time"
)

func Create(postData postEntities.Post, db *sql.DB) (post postEntities.Post, err error) {
	query := `INSERT INTO posts (title, content, created_at) VALUES ($1, $2, $3) RETURNING id`

	result, err := db.Exec(query, postData.Title, postData.Content, time.Now())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RESULT", result)
	return postEntities.Post{}, err
}
