package router

import (
	"fmt"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "internal/app/assets/templates/*.gohtml",
	}

	rnd = renderer.New(opts)
}

func SetUpRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOME")
	err := rnd.HTML(w, http.StatusOK, "home", nil)
	if err != nil {
		fmt.Printf("Ошибка рендеринга шаблона: %v\n", err)
		http.Error(w, "Ошибка рендеринга", http.StatusInternalServerError)
		return
	}
}
