package control

import (
	"fmt"
	"io"
	"net/http"
)

// AddArticle AddArticle
func AddArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("title")
	r.ParseForm()
	title := r.Form.Get("title")
	fmt.Println(title)
	io.WriteString(w, title)
}
