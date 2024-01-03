package main

import (
	"fmt"
	"math/rand"
)

func main() {

	unique := map[int]bool{}
	var slice []int
	for i := 0; i < 10; i++ {
		t := rand.Intn(10)
		if !unique[t] {
			unique[t] = true
			slice = append(slice, t)
		} else {
			i--
		}

	}
	
	fmt.Println(slice)
}
