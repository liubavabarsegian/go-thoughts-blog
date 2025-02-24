package handlers

import (
	"day06/internal/app/service/post"
	"encoding/json"
	"net/http"
)

// PostHandler обрабатывает запросы, связанные с постами
type PostHandler struct {
	postService post.Service
}

// NewPostHandler создает новый хендлер для постов
func NewPostHandler(postService post.Service) *PostHandler {
	return &PostHandler{postService: postService}
}

// GetAllPosts — обработчик получения всех постов
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
