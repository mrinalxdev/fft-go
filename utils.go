package main

import (
	"math"
	"math/cmplx"
)

func generateSignal(size int) []complex128 {
	signal := make([]complex128, size)
	for i := 0; i < size; i ++ {
		signal[i] = complex(math.Sin(2*math.Pi*float64(i) / float64(size)), 0)
	}
	return signal
}

func magnitude(c complex128) float64 {
	return cmplx.Abs(c)
}
