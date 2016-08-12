package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Lowef() {
	start := time.Now()
	s, sep := "", ""
	for i := 0; i < 10000; i++ {
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}

		fmt.Println(s)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func Highef() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Print(strings.Join(os.Args[1:], " "))
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	Lowef()
	Highef()
}

//进入chap1/ex1.3  go build以后 ./ch1 ex1.3 arg1得到10000次循环Lowef()花费13.67s,Highef()花费0.01s
