package handlers

import (
	"day06/internal/app/models"
	"day06/internal/app/service/post"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PostHandler обрабатывает запросы, связанные с постами
type PostHandler struct {
	postService post.Service
}

type GetPostResponse struct {
	Post        models.Post
	CurrentPage uint
}
type GetAllPostsResponse struct {
	Posts        []models.Post
	HasPrevious  bool
	HasNext      bool
	CurrentPage  uint
	PreviousPage uint
	NextPage     uint
}

// NewPostHandler создает новый хендлер для постов
func NewPostHandler(postService post.Service) *PostHandler {
	return &PostHandler{postService: postService}
}

// GetPost - получить пост по id
func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Failed to parse post id", http.StatusInternalServerError)
		return
	}

	post, err := h.postService.GetPost(uint(id))
	if err != nil {
		http.Error(w, "Failed to get the post", http.StatusInternalServerError)
		return
	}

	parsed_page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	page := uint(parsed_page)

	response := &GetPostResponse{
		Post:        post,
		CurrentPage: page,
	}

	tmpl := template.Must(template.ParseFiles("/app/internal/templates/show.gohtml"))
	tmpl.Execute(w, response)
}

// GetAllPosts — обработчик получения всех постов
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	parsed_page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || parsed_page < 1 {
		parsed_page = 1
	}

	page := uint(parsed_page)
	posts, err := h.postService.GetAllPosts(page)
	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}

	response := &GetAllPostsResponse{
		Posts:        posts,
		HasPrevious:  page > 1,
		HasNext:      page < uint(len(posts)),
		CurrentPage:  page,
		PreviousPage: page - 1,
		NextPage:     page + 1,
	}

	tmpl := template.Must(template.ParseFiles("/app/internal/templates/index.gohtml"))
	tmpl.Execute(w, response)
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
