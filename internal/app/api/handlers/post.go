package handlers

import (
	"day06/internal/app/models"
	"day06/internal/app/service/post"
	"encoding/json"
	"fmt"
	"html/template"
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

	tmpl := template.Must(template.ParseFiles("/app/internal/templates/index.gohtml"))
	tmpl.Execute(w, posts)
}

// CreatePost — обработчик создания поста
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE POST")
	var postModel models.Post
	json.NewDecoder(r.Body).Decode(&postModel)
	post, err := h.postService.CreatePost(postModel)
	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
