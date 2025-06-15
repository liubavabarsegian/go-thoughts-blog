package login

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

// LoginService реализует интерфейс Service
type LoginService struct {
}

// NewLoginService создает новый сервис для логина
func NewService() Service {
	return &LoginService{}
}

func (s *LoginService) IsLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("admin_auth")
	return err == nil && cookie.Value == "true"
}

func (s *LoginService) Login(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  "admin_auth",
		Value: "true",
		Path:  "/",
	})
}

func (s *LoginService) Logout(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "admin_auth",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
}

func (s *LoginService) GetAdminCreds() (string, string) {
	file, err := os.Open("/app/.admin_credentials.txt")
	if err != nil {
		return "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		}
	}

	return "", ""
}
