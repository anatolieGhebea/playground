package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	fib := fibonacci(6)
	if fib != 8 {
		t.Errorf("Fibonacci(6) = %d; want 8", fib)
	}
}

func TestFibonacciMemo(t *testing.T) {
	fib := fibonacciMemo(6, make(map[int]int64))
	if fib != 8 {
		t.Errorf("FibonacciMemo(6) = %d; want 8", fib)
	}
}
