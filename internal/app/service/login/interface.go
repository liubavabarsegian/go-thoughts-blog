package login

import "net/http"

// Service определяет методы, которые должен реализовывать сервис логина
type Service interface {
	// to do: change to cookies in args
	IsLoggedIn(r *http.Request) bool
	Login(w http.ResponseWriter)
	Logout(w http.ResponseWriter)
	GetAdminCreds() (string, string)
}
