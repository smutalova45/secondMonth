package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	time.Now()
	wg.Add(2)

	go task1(&wg)
	go task2(&wg)
	
	wg.Wait()
	fmt.Println(time.Since(time.Now()))

}
func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 65; i < 72; i++ {
		fmt.Printf("%c\n", i)
	}

}
func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 13; i++ {
		fmt.Println(i)
	}
}
