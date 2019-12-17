package router

import (
	"github.com/01super/blog_web/control"
	"net/http"
)

// Run 连接数据库
func Run() {
	http.HandleFunc("/api/addArticle", control.AddArticle)
	// http.HandleFunc("/", control.IndexView)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":9000", nil)
}
