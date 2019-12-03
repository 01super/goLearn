package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func countStringChinese() {
	str := "Hello Go 语言"
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)
}

func countEnWorld() {
	var str string = "Hello world, go come go world"
	s1 := strings.Split(str, " ")
	m := make(map[string]int, 10)
	for _, value := range s1 {
		if key, ok := m[value]; !ok {
			m[value] = 1
			fmt.Println("key", key)
		} else {
			m[value]++
		}
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// 可变参数 ...  fs("hello", 1, 2, 3, 4, 5)
func fs(x string, y ...int) int {
	fmt.Println(y) // y是一个int类型的切片[1 2 3 4 5]
	return y[0]
}

// 命名返回值
func nameed(x, y int) (sum int) {
	sum = x + y // 如果使用命名的返回值，在函数中可以直接使用返回值变量
	return      //如果使用命名的返回值，return后面可以省略返回值变量
}

func main() {
	// countStringChinese()
	// countEnWorld()
	// fs("hello", 1, 2, 3, 4, 5)
	fmt.Println(nameed(1, 2))
}
