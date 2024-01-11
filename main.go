package main

import "fmt"

func main() {

	v := writeWords()
	ch := removeRepeated(v)

	for v := range ch {
		fmt.Println(v)
	}
}

func writeWords() chan string {
	slice := []string{"apple", "cherry", "banana", "cherry","apple"}
	ch1 := make(chan string)

	go func() {
		defer close(ch1)
		for _, word := range slice {
			ch1 <- word
		}
	}()

	return ch1
}

func removeRepeated(ch chan string) chan string {

	ch2 := make(chan string)
	go func() {
		defer close(ch2)
		m := map[string]bool{}
		for v := range ch {
			if !m[v] {
				m[v] = true
				ch2 <- v

			}
		}
	}()
	return ch2
}
