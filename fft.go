package main

import (
	"math"
	"math/cmplx"
	"sync"
)

func ParallelFFT(input []complex128) []complex128 {
	n := len(input)

	if n <= 1 {
		return input
	}

	// spliting the input into even and
	// odd incdices

	// this splits the data and
	// fiddle factor
	// w = z (5 * 4)

	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)

	for i := 0; i < n/2; i++ {
		even[i] = input[i*2]
		odd[i] = input[i*2+1]
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var evenFFT, oddFFT []complex128

	go func() {
		defer wg.Done()
		evenFFT = ParallelFFT(even)
	}()

	go func() {
		defer wg.Done()
		oddFFT = ParallelFFT(odd)
	}()

	wg.Wait()

	output := make([]complex128, n)
	for k := 0; k < n/2; k++ {
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(n))) * oddFFT[k]
		output[k] = evenFFT[k] + t
		output[k+n/2] = evenFFT[k] - t
	}

	return output

}
