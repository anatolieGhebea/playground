package main

import (
	"fmt"
	"time"
)

type FibResult struct {
	n                  int
	result             int64
	resultMemo         int64
	resultMemoShared   int64
	duration           time.Duration
	durationMemo       time.Duration
	durationMemoShared time.Duration
}

var memo = map[int]int64{}

func main() {
	// ask user for a number
	fmt.Println("1. Calculate Fibonacci")
	fmt.Println("2. Calculate Fibonacci 2")
	fmt.Print("Enter the name of the function to run: ")
	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		calculateFib()
	case 2:
		fibExecStats()
	}
}

func fibExecStats() {
	numbers := 50
	fn_start := time.Now()
	results := make([]FibResult, numbers)

	fmt.Println("\nCalculating Fibonacci numbers from 1 to 50...")
	fmt.Println("\nNumber\tResult\t\tTime Taken Fib\t\tTime Taken Memo\t\tTime Taken Memo Shared")
	fmt.Println("----------------------------------------------------------------")

	nfrom := 1
	for i := nfrom; i <= nfrom+(numbers-1); i++ {
		start := time.Now()
		result := fibonacci(i)
		duration := time.Since(start)

		start = time.Now()
		resultMemo := fibonacciMemo(i, make(map[int]int64))
		durationMemo := time.Since(start)

		start = time.Now()
		resultMemoShared := fibonacciMemo(i, memo)
		durationMemoShared := time.Since(start)

		results[i-nfrom] = FibResult{
			n:                  i,
			result:             result,
			resultMemo:         resultMemo,
			resultMemoShared:   resultMemoShared,
			duration:           duration,
			durationMemo:       durationMemo,
			durationMemoShared: durationMemoShared,
		}

		// Convert duration to seconds (as float64)
		seconds := duration.Seconds()
		secondsMemo := durationMemo.Seconds()
		secondsMemoShared := durationMemoShared.Seconds()
		fmt.Printf("%-7d\t%-10d\t%.6f\t\t%.6f\t\t%.6f\n", i, resultMemo, seconds, secondsMemo, secondsMemoShared)
	}

	fn_end := time.Since(fn_start)
	fmt.Printf("\nTotal time taken: %s\n", fn_end)
}

func calculateFib() {
	// ask user for a number
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanln(&input)

	// track time
	start := time.Now()

	// calculate fibonacci
	result := fibonacci(input)

	// calculate elapsed time
	elapsed := time.Since(start)

	fmt.Printf("Fibonacci(%d) = %d\n", input, result)
	fmt.Printf("Time taken: %s\n", elapsed)
}

func fibonacci(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciMemo(n int, memo map[int]int64) int64 {
	// Base cases
	if n <= 1 {
		return int64(n)
	}

	// Check if we've already calculated this value
	if val, exists := memo[n]; exists {
		return val
	}

	// Calculate and store the result
	memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	return memo[n]
}
