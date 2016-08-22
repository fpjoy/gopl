package main

import (
	"crypto/sha256"
	"fmt"
)

//支持传入string型
func ShaCount(x, y string) int {
	var count int
	shaX := sha256.Sum256([]byte(x))
	shaY := sha256.Sum256([]byte(y))
	length := len(shaX)

	for i := 0; i < length; i++ {
		if shaX[i] != shaY[i] {
			count++
		}
	}
	return count

}
