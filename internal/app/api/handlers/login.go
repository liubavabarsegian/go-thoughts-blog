package handlers

import (
	"day06/internal/app/service/login"
	"html/template"
	"net/http"
)

// LoginHandler обрабатывает запросы, связанные с авторизацией
type LoginHandler struct {
	loginService login.Service
}

// NewLoginHandler создает новый хендлер для авторизацииы
func NewLoginHandler(loginService login.Service) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

func (h *LoginHandler) AdminPage(w http.ResponseWriter, r *http.Request) {
	if !h.loginService.IsLoggedIn(r) {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("/app/internal/templates/admin.gohtml"))
	tmpl.Execute(w, nil)
}

func (h *LoginHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("/app/internal/templates/login.gohtml"))
	tmpl.Execute(w, nil)
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	adminUser, adminPass := h.loginService.GetAdminCreds()
	if username == adminUser && password == adminPass {
		h.loginService.Login(w)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("/app/internal/templates/login.gohtml"))
	tmpl.Execute(w, nil)
}

func (h *LoginHandler) Logout(w http.ResponseWriter, r *http.Request) {
	h.loginService.Logout(w)
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
