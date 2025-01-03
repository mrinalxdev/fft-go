package main


func ParallelFFT(input []complex128) []complex128 {
	n := len(input)

	if n <= 1 {
		return input
	}

	//spliting the input into even and
	// odd incdices

	even := make([]complex128, n/2)
	odd := make([]complex128, n/2)

	


}
