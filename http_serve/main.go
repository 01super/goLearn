package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(res http.ResponseWriter, req *http.Request) {
	index, err := ioutil.ReadFile("./dist/index.html")
	if err != nil {
		fmt.Println("文件读取失败：", err)
	} else {
		res.Write(index)
	}
}
func main() {
	http.HandleFunc("/home", f1)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
