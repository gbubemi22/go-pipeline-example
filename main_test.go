package main

import (
	"testing"
)

// TestGenerator tests the Generator function.
func TestGenerator(t *testing.T) {
	out := make(chan int, 10)
	go Generator(10, out)

	count := 0
	for num := range out {
		if num < 1 || num > 100 {
			t.Errorf("Generated number out of bounds: %d", num)
		}
		count++
	}
	if count != 10 {
		t.Errorf("Expected 10 numbers, got %d", count)
	}
}

// TestSquare tests the Square function.
func TestSquare(t *testing.T) {
	in := make(chan int, 10)
	out := make(chan int, 10)

	go func() {
		for _, v := range []int{1, 2, 3, 4, 5} {
			in <- v
		}
		close(in)
	}()

	go Square(in, out)

	expected := []int{1, 4, 9, 16, 25}
	for i, expectedValue := range expected {
		got := <-out
		if got != expectedValue {
			t.Errorf("Square of %d = %d; want %d", i+1, got, expectedValue)
		}
	}
}

// TestSum tests the Sum function.
func TestSum(t *testing.T) {
	in := make(chan int, 5)

	go func() {
		for _, v := range []int{1, 4, 9, 16, 25} {
			in <- v
		}
		close(in)
	}()

	got := Sum(in)
	expected := 55 // 1 + 4 + 9 + 16 + 25
	if got != expected {
		t.Errorf("Sum = %d; want %d", got, expected)
	}
}
