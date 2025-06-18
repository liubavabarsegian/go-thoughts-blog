package handlers

import (
	"day06/internal/app/models"
	"day06/internal/app/service/post"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
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
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	if page < 1 {
		page = 1
	}

	posts, err := h.postService.GetAllPosts(page)
	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("/app/internal/templates/index.gohtml"))
	tmpl.Execute(w, posts)
}

// CreatePost — обработчик создания поста
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
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
