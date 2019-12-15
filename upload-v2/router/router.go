package router

import (
	"fmt"
	"github.com/01super/upload-v2/control"
	"github.com/01super/upload-v2/model"
	"net/http"
)

// Run 连接数据库
func Run() {
	http.HandleFunc("/", control.IndexView)
	http.HandleFunc("/upload", control.UploadView)
	http.HandleFunc("/api/upload", control.APIUpload)
	http.HandleFunc("/detial", control.DetialView)
	http.HandleFunc("/api/list", control.APIList)
	http.HandleFunc("/list", control.ListVIew)
	http.HandleFunc("/api/delete", control.APIDrop)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":8080", nil)
	mod := model.Info{}
	err := model.Db.Get(&mod, "select * from info where id=1")
	fmt.Println(mod, err)
}
