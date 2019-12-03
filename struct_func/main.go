package main

import "fmt"

type dog struct {
	name string
}

func newDog(name string) *dog {
	return &dog{
		name,
	}
}
func (d dog) wang() {
	fmt.Println(d.name + "wangwangwang")
}

type myInt int

func (m myInt) sayNumber(num int) {
	fmt.Println(num)
}
func main() {
	d1 := newDog("阿黄")
	d1.wang()
	m := myInt(100)
	m.sayNumber(666)
}
