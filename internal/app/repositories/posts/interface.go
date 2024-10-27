package posts

import postEntities "day06/internal/app/entities/post"

type PostRepository interface {
	GetByID(id int) (*postEntities.Post, error)
	GetAll() ([]*postEntities.Post, error)
	Create(post *postEntities.Post) error
	Update(post *postEntities.Post) error
	Delete(id int) error
}
