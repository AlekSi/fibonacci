package services

import (
	"github.com/AlekSi/fibonacci/errs"
)

// Fibonacci service computes and stores Fibonacci numbers.
// It's expected to be long-lived (one per process).
type Fibonacci struct {
	numbers []uint
}

// NewFibonacci creates new Fibonacci service instance.
// It precomputes and stores Fibonacci numbers up to maximum uint.
func NewFibonacci() *Fibonacci {
	numbers := make([]uint, 0, 100)
	numbers = append(numbers, 0, 1)
	for i := 2; i < 95; i++ {
		numbers = append(numbers, numbers[i-2]+numbers[i-1])
	}
	return &Fibonacci{
		numbers: numbers,
	}
}

// Numbers returns first n Fibonacci numbers, or APIError if n is too big.
// This method is thread-safe.
func (f *Fibonacci) Numbers(n uint) ([]uint, error) {
	if uint(len(f.numbers)) <= n {
		return nil, errs.New(errs.InvalidParameter, "n is too big", nil)
	}
	return f.numbers[:n], nil
}
