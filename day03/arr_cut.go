package main

import "fmt"

func main() {
	var arr1 [3]int
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr), cap(arr), arr1[0])
	for idx, v := range arr1 {
		fmt.Println(idx, v)
	}
	slice1 := make([]int, 3, 5)
	var arr2 = append(slice1, 1, 2, 3)
	arr3 := arr[:]
	fmt.Println(arr2, len(arr2), arr3)
}
