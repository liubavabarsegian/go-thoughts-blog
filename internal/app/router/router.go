package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

// func init() {
// 	opts := renderer.Options{
// 		ParseGlobPattern: "internal/app/assets/templates/*.gohtml",
// 	}

//		rnd = renderer.New(opts)
//	}
func SetupHandlers(mux *mux.Router) {
	mux.HandleFunc("/", home)
	// mux.HandleFunc("/admin", adminHandler)
}

func home(w http.ResponseWriter, r *http.Request) {
	// logger := log.Ctx(r.Context())
	// page := r.URL.Query().Get("page")
	// articles, hasNext, hasPrevious := fetchArticles(page)

	// data := struct {
	// 	Posts        []postEntities.Post
	// 	HasNext      bool
	// 	HasPrevious  bool
	// 	NextPage     int
	// 	PreviousPage int
	// }{
	// 	Posts: post.GetAll(),
	// }

	// logger.Println("INFOO")
	err := rnd.HTML(w, http.StatusOK, "home", nil)
	if err != nil {
		fmt.Printf("Ошибка рендеринга шаблона: %v\n", err)
		http.Error(w, "Ошибка рендеринга", http.StatusInternalServerError)
		return
	}
}
