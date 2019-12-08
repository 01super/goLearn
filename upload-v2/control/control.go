package control

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// Route 路由
func IndexView(w http.ResponseWriter, r *http.Request) {
	html := loadHtml("../views/index.html")
	w.Write(html)

}

func UploadView(w http.ResponseWriter, r *http.Request) {
	html := loadHtml("../views/upload.html")
	w.Write(html)
}

func ApiUpload(w http.ResponseWriter, r *http.Request) {
	f, h, err := r.FormFile("file")
	if err != nil {
		io.WriteString(w, "上传错误")
		return
	}
	t := h.Header.Get("Content-Type")
	if !strings.Contains(t, "image") {
		io.WriteString(w, "上传格式错误")
		return
	}
	os.Mkdir("./static")
	name := time.Now().Format("2016-0102150405") + path.Ext(h.Filename) // get .xxx
	out, err := os.Create("./static/" + name)
	if err != nil {
		io.WriteString(w, "文件创建失败")
		return
	}
	io.Copy(out, f)
	defer out.Close()
}

func loadHtml(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("load file failed666")
		return []byte("")
	}
	return buf
}
