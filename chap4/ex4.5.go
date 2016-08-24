package main

import "fmt"

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func moveDup(s []string) []string {

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			s = remove(s, i)
			i--
		}

	}
	return s
}

func main() {
	s := []string{"hi", "hello", "hello", "hi", "hi", "hi", "hi", "hello", "hx", "hx"}
	fmt.Println(moveDup(s))
}
