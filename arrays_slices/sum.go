package arrays_slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	numSlices := len(slicesToSum)         // num slices that were passed in as args
	sliceSum := make([]int, 0, numSlices) // makes a slice to store the sums, one slot per slice, length of 0

	for _, numbers := range slicesToSum { // iterate over a slice
		sliceSum = append(sliceSum, Sum(numbers)) // appends the sum to the end of the initially len 0 slice
	}

	return sliceSum
}

func SumAllTails(slicesToSum ...[]int) []int {
	numSlices := len(slicesToSum)         // num slices that were passed in as args
	sliceSum := make([]int, 0, numSlices) // makes a slice to store the sums, one slot per slice, length of 0

	for _, numbers := range slicesToSum { // iterate over a slice
		if len(numbers) == 0 {
			sliceSum = append(sliceSum, 0)
		} else {
			tail := numbers[1:]
			sliceSum = append(sliceSum, Sum(tail)) // appends the sum to the end of the initially len 0 slice
		}
	}

	return sliceSum
}
