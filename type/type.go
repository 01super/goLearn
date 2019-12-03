package main

import "fmt"

type myInt int  // 自定义类型
type yInt = int //类型别名

func main() {
	var a myInt = 10
	fmt.Printf("%T\n", a) // main.myInt
	var b yInt = 20
	fmt.Printf("%T\n", b) // int
	var s struct {        //匿名结构体
		name string
		age  int
	}
	s.name = "xiaoming"
	s.age = 18
	fmt.Printf("%T\n%v", s, s)
	type ty struct {
		name int
	}
	var ss = new(ty)
	ss.name = 12
	p := &ss
	fmt.Printf("%p", ss)
	fmt.Println(p)
}
