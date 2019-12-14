package control

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/01super/upload-v2/model"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// IndexView send index.html
func IndexView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("../views/index.html")
	w.Write(html)

}

// UploadView upload.html
func UploadView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("../views/upload.html")
	w.Write(html)
}

// APIUpload recive file
func APIUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // 处理数据，为了后面能拿到请求中的备注信息
	f, h, err := r.FormFile("file") // FormFile只会处理文件相关的数据，不会处理字符类的
	defer f.Close()
	if err != nil {
		io.WriteString(w, "上传错误")
		return
	}
	t := h.Header.Get("Content-Type")
	if !strings.Contains(t, "image") {
		io.WriteString(w, "上传格式错误")
		return
	}
	os.Mkdir("./static", 0666)
	name := time.Now().Format("2016-0102150405") + path.Ext(h.Filename) // get .xxx
	fmt.Println(name)
	out, err := os.Create("./static/" + name)
	if err != nil {
		io.WriteString(w, "文件创建失败"+err.Error())
		return
	}
	_, err3 := io.Copy(out, f)
	if err3 != nil {
		io.WriteString(w, "文件保存失败，"+err3.Error())
	}
	defer out.Close()
	mod := model.Info{
		Name: h.Filename,
		Path: "/static/" + name,
		Unix: time.Now().Unix(),
		Note: r.Form.Get("note"),
	}
	err2 := model.InfoAdd(&mod)
	if err2 != nil {
		fmt.Println("插入数据失败")
	}
	http.Redirect(w, r, "list", 302)
}

// DetialView 详细页面
func DetialView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	mod, _ := model.InfoGet(id)
	html := loadHTML("../views/detial.html")
	date := time.Unix(mod.Unix, 0).Format("2006年02月12日15:01:02")
	html = bytes.Replace(html, []byte("@src"), []byte(mod.Path), 1)
	html = bytes.Replace(html, []byte("@unix"), []byte(date), 1)
	html = bytes.Replace(html, []byte("@note"), []byte(mod.Note), 1)
	w.Write(html)
}

// APIList 获取列表
func APIList(w http.ResponseWriter, r *http.Request) {
	mod, err := model.InfoList()
	if err != nil {
		w.Write([]byte("获取失败"))
	}
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-type", "application/json")
	w.Write(buf)
}

// ListVIew listview
func ListVIew(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("../views/list.html")
	w.Write(html)
}

func loadHTML(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("load file failed666")
		return []byte("")
	}
	return buf
}
