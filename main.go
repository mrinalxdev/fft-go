package main

import (
	"fmt"
	"time"
)

func main(){
	// signal size
	n := 1024
	signal := generateSignal(n)

	start := time.Now()
	result := ParallelFFT(signal)
	elapsed := time.Since(start)

	fmt.Printf("FFT completed in %v\n", elapsed)
	fmt.Println("First 10 FFT results : ")
	for i := 0; i < 10; i++ {
		fmt.Printf("Result[%d]: %v\n", i, result[i])
	}
}
