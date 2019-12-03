package main

import "fmt"

func f1() int { //没有命名的函数的返回值
	x := 5
	defer func() {
		x++
	}()
	return x //5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值是x，先将x赋值为5，后将x++，返回值变为6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值是y，将x赋值给y，x++，不影响y，返回5
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //返回值为x，将x赋值为5，defer函数加的是拷贝的x的值，返回值仍为5
}

func main() {
	fmt.Println(f1())
}
