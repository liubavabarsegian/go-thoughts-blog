package router

import (
	"day06/internal/app/api/handlers"
	"day06/internal/app/service"
	"day06/internal/app/service/login"
	"day06/internal/app/service/post"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupHandlers настраивает маршруты и передает сервисы в хендлеры
func SetupHandlers(router *mux.Router, services *service.Service) {
	setuploginHandler(router, services.Login)
	setupPostHandler(router, services.Post)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}

func setupPostHandler(router *mux.Router, postService post.Service) {
	postHandler := handlers.NewPostHandler(postService)

	router.HandleFunc("/", postHandler.GetAllPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts/create", postHandler.CreatePost).Methods(http.MethodPost)
	router.HandleFunc("/posts/{id}", postHandler.GetPost).Methods(http.MethodGet)
}

func setuploginHandler(router *mux.Router, loginService login.Service) {
	loginHandler := handlers.NewLoginHandler(loginService)

	router.HandleFunc("/admin", loginHandler.AdminPage).Methods(http.MethodGet)
	router.HandleFunc("/admin/login", loginHandler.LoginPage).Methods(http.MethodGet)
	router.HandleFunc("/admin/login/submit", loginHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/admin/logout", loginHandler.Logout).Methods(http.MethodGet)
}
