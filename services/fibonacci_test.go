package services

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlekSi/fibonacci/errs"
)

func TestFibonacci(t *testing.T) {
	var wg sync.WaitGroup
	f := NewFibonacci()
	for n, expected := range map[uint][]uint{
		0: {},
		1: {0},
		2: {0, 1},
		3: {0, 1, 1},
		4: {0, 1, 1, 2},
		5: {0, 1, 1, 2, 3},
		39: {0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946,
			17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887,
			9227465, 14930352, 24157817, 39088169}, // https://oeis.org/A000045
	} {
		wg.Add(1)
		go func(n uint, expected []uint) {
			num, err := f.Numbers(n)
			assert.NoError(t, err)
			assert.Equal(t, expected, num)
			wg.Done()
		}(n, expected)
	}

	for n, expected := range map[uint]uint{
		91: 2880067194370816120,
		92: 4660046610375530309,
		93: 7540113804746346429,
		94: 12200160415121876738,
	} {
		wg.Add(1)
		go func(n uint, expected uint) {
			num, err := f.Numbers(n)
			assert.NoError(t, err)
			assert.Equal(t, expected, num[len(num)-1])
			wg.Done()
		}(n, expected)
	}

	num, err := f.Numbers(95)
	assert.Zero(t, num)
	assert.Equal(t, errs.New(errs.InvalidParameter, "n is too big", nil), err)

	wg.Wait()
}
