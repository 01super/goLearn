package main

import "net/http"

import "fmt"

import "io/ioutil"

import "os"

import "io"

import "bytes"

import "log"

import "strings"

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		html := loadHTML("./upload.html")
		w.Write(html)
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			w.Write([]byte("文件上传有误" + err.Error()))
			return
		}
		t := h.Header.Get("Content-Type")
		if !strings.Contains(t, "image") {
			io.WriteString(w, "<html><a href=\"/upload\">请上传图片</a></html")
			return
		}
		os.Mkdir("./static", 0666) // 以main.go作为相对路径
		out, err := os.Create("./static/" + h.Filename)
		if err != nil {
			io.WriteString(w, "文件创建失败："+err.Error())
			return
		}
		_, err1 := io.Copy(out, f)
		if err1 != nil {
			io.WriteString(w, "文件保存失败："+err1.Error())
		}
		// io.WriteString(w, "static/"+h.Filename) // web服务器，相当于main.go
		http.Redirect(w, r, "/img?name="+h.Filename, 302)
		return
	}
}

func imageView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	f, err := os.Open("./static/" + name)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	defer f.Close()
	w.Header().Set("Content-Type", "image")
	io.Copy(w, f)
}

func image(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	html := loadHTML("./img.html")
	html = bytes.Replace(html, []byte("@src"), []byte("/image?name="+name), 1)
	w.Write(html)
}

func listView(w http.ResponseWriter, r *http.Request) {
	names, err := ioutil.ReadDir("./static/")
	if err != nil {
		log.Println("文件读取失败：", err)
		return
	}
	html := loadHTML("./list.html")
	str := ""
	for _, v := range names {
		str += `<li><img src="/image?name=` + v.Name() + `" alt="未发现"></li>`
	}
	html = bytes.Replace(html, []byte("@list"), []byte(str), 1)
	w.Write(html)
}

func homeView(w http.ResponseWriter, r *http.Request) {
	html := loadHTML("./index.html")
	w.Write(html)
}

func loadHTML(path string) []byte {
	const errStr string = "<html><head><body><h3>Errors</h3></body></head></html>"
	f, err := os.Open(path)
	if err != nil {
		return []byte(errStr)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte(errStr)
	}
	return buf
}

func main() {
	http.HandleFunc("/", homeView)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/image", imageView)
	http.HandleFunc("/img", image)
	http.HandleFunc("/list", listView)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务启动失败：", err)
		return
	}
	fmt.Println("服务启动成功... ...")
}
