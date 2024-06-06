package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// worker function to calculate the sum of a part of the array
func worker(arr []int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, num := range arr {
		sum += num
	}
	resultChan <- sum
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const arraySize = 1000000
	arr := make([]int, arraySize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	// Number of goroutines to use
	const numWorkers = 10

	// Channel to collect results from goroutines
	resultChan := make(chan int, numWorkers)

	// WaitGroup to ensure all goroutines complete before proceeding
	var wg sync.WaitGroup

	// Calculate the size of each chunk to be processed by each goroutine
	chunkSize := arraySize / numWorkers

	// Start goroutines
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = arraySize // Last chunk
		}

		wg.Add(1)
		go worker(arr[start:end], resultChan, &wg)
	}

	// Close the result channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Aggregate results from all goroutines
	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Printf("The total sum of the array elements is: %d\n", totalSum)
}
