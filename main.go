package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)



// Generate Random number

func Generator(count int, out chan<- int) {
	defer close(out) // Need to close the channel after sending all numbers
	 for i := 0; i < count; i++ {
		out <- rand.Intn(100) + 1
	 }
}


//Sequre the numbers from Generator
func Square(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out) // Close channel when done.
}


// Sum receives squared numbers and computes their sum.
func Sum(in <-chan int) int {
	total := 0
	for num := range in {
		total += num
	}
	return total
}

func main() {
	count := flag.Int("count", 10000, "Number of random integers to generate")
	flag.Parse()

	

	// Create channels
	genChan := make(chan int, 100)
	squareChan := make(chan int, 100)

	startTime := time.Now() // Start time for processing

	// Using WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(3)


	// Start the generator goroutine
	go func() {
		defer wg.Done()
		Generator(*count, genChan)
	}()

	// Start the squaring goroutine
	go func() {
		defer wg.Done()
		Square(genChan, squareChan)
	}()

	// Start the sum calculation
	go func() {
		defer wg.Done()
		sum := Sum(squareChan)
		fmt.Printf("Sum of squares: %d\n", sum)
	}()

	wg.Wait() // Wait for all goroutines to finish
	endTime := time.Since(startTime)
	fmt.Printf("Processing took %v\n", endTime)

}

