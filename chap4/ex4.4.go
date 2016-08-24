package main

import "fmt"

// s 准备旋转的整型数组切片
// n 向左旋转n个元素
func Rotate(s []int, n int) []int {
	//将s复制一份拼接到原来的s切片上
	length := len(s)
	for _, r := range s {
		s = append(s, r)
	}

	n = n % length
	rotates := s[n : length+n]
	return rotates
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(Rotate(s, 2))
	fmt.Println(Rotate(s, 8))
}
