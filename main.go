package main

import (
	"fmt"
	"time"
)

func main(){
	// signal size
	n := 1024
	signal := generateSignal(n)

	if err := plotSignal(signal, "input_signal.png"); err != nil {
		fmt.Printf("Error Plotting signal : %v\n", err)
		return
	}

	fmt.Println("Input signal plot saved as 'input_signal.png'")


	start := time.Now()
	result := ParallelFFT(signal)
	elapsed := time.Since(start)

	if err := plotFFT(result, "fft_magnitude.png"); err != nil {
		fmt.Printf("Error plotting FFT: %v\n", err)
		return
	}
	fmt.Println("FFT magnitude spectrum plot saved as 'fft_magnitude.png'")
	fmt.Printf("FFT completed in %v\n", elapsed)
}
