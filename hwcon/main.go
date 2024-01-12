package main

import (
	"fmt"
	"sort"
	"sync"
)

func exercise1() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 1; i <= 19; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		for num := range ch {
			fmt.Println(num)
		}
	}()

	wg.Wait()
}

func exercise2() {
	ch := make(chan int)
	var wg sync.WaitGroup
	sum := 0

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		sum += num
	}

	fmt.Println("Sum:", sum)
}

func exercise3(n int) {
	ch := make(chan int)
	var wg sync.WaitGroup
	result := 1

	wg.Add(1)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		for num := range ch {
			result *= num
		}
	}()

	wg.Wait()
	fmt.Println("Factorial of", n, "is", result)
}
func exercise9(numbers []int) {
	ch := make(chan []int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sort.Ints(numbers)
		ch <- numbers
		close(ch)
	}()

	go func() {
		wg.Wait()
	}()

	sortedNumbers := <-ch
	fmt.Println("Sorted numbers:", sortedNumbers)
}

func exercise10() {
	ch := make(chan int)
	var wg sync.WaitGroup
	sum := 0

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n * n
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		sum += num
	}

	fmt.Println("Sum of squares:", sum)
}

func exercise11(n int) {
	ch := make(chan int)
	var wg sync.WaitGroup
	fibonacci := []int{0, 1}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; fibonacci[i-1]+fibonacci[i-2] <= n; i++ {
			fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
		}
		ch <- len(fibonacci)
		close(ch)
	}()

	go func() {
		wg.Wait()
	}()

	fmt.Println("Fibonacci sequence length:", <-ch)
}

func exercise12(numbers []int) {
	ch := make(chan int)
	var wg sync.WaitGroup
	sum := 0

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if n%2 == 0 {
				ch <- n
			}
		}(num)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		sum += num
	}

	fmt.Println("Sum of even numbers:", sum)
}

func exercise13(input string) {
	ch := make(chan int)
	var wg sync.WaitGroup
	longestWord := ""

	wg.Add(1)
	go func() {
		defer wg.Done()
		words := splitWords(input)
		maxLength := 0
		for _, word := range words {
			if len(word) > maxLength {
				maxLength = len(word)
				longestWord = word
			}
		}
		ch <- maxLength
		close(ch)
	}()

	go func() {
		wg.Wait()
	}()

	fmt.Println("Length of the longest word:", <-ch, "Longest word:", longestWord)
}

func splitWords(input string) []string {

	return []string{"example", "words", "from", "input"}
}

func exercise14(files []string) {
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			linesCount := processFile(filename)
			ch <- linesCount
		}(file)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalLines := 0
	for count := range ch {
		totalLines += count
	}

	fmt.Println("Total lines in all files:", totalLines)
}

func processFile(filename string) int {
	// Process the file and return the count of lines
	return 10 // Example: Count of lines in a file
}

func exercise15() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

func main() {
	fmt.Println("Exercise 1:")
	exercise1()

	fmt.Println("\nExercise 2:")
	exercise2()

	fmt.Println("\nExercise 3:")
	exercise3(5)
	fmt.Println("\nExercise 9:")
	nums:=[]int{2,3,4,5}
	exercise9(nums)

	fmt.Println("\nExercise 10:")
	exercise10()

	fmt.Println("\nExercise 11:") // Example: Calculate Fibonacci sequence up to 50
	exercise11(50)

	fmt.Println("\nExercise 12:")
	exercise12([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) // Example: Sum of even numbers in a slice

	fmt.Println("\nExercise 13:")
	exercise13("This is an example input string") // Example: Find the length of the longest word in a string

	fmt.Println("\nExercise 14:")
	exercise14([]string{"file1.txt", "file2.txt", "file3.txt"}) // Example: Concurrent file processing

	fmt.Println("\nExercise 15:")
	exercise15()
}
