package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/01super/blog_web/model"
)

// AddArticle AddArticle
func AddArticle(w http.ResponseWriter, r *http.Request) {
	// 检查是否为post请求
	fmt.Println("http.MethodPost", http.MethodPost)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}
	//application/x-www-form-urlencoded 格式
	//r.ParseForm()
	//name:=r.PostFormValue("name")
	//fmt.Fprintf(w, "x-www-form-urlencoded -> name="+name+"\n")

	//multipart/form-data 格式
	//r.ParseMultipartForm(32<<20)
	//name:=r.PostFormValue("name")
	//fmt.Fprintf(w, "multipart/form-data -> name="+name+"\n")

	//application/json格式
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body) //获取post的数据
	fmt.Println(con)

	mod := model.Artical{}
	fmt.Println(mod)
	json.Unmarshal(con, &mod)
	fmt.Println(mod.Title)
	model.ArticalAdd(&mod)
}

// ListInfo get blog list
func ListInfo(w http.ResponseWriter, r *http.Request) {
	mod, err := model.ArticalList()
	if err != nil {
		w.Write([]byte("查询列表失败"))
	}
	buf, err := json.Marshal(mod)
	w.Write(buf)
}

// ArticleDetial detial
func ArticleDetial(w http.ResponseWriter, r *http.Request) {
	fmt.Println("im in")
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	mod, err := model.ArticalGet(id)
	if err != nil {
		w.Write([]byte("获取失败"))
	}
	buf, _ := json.Marshal(mod)
	w.Write(buf)
}
