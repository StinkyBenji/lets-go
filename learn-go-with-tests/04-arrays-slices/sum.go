package main

// [N/...]type{value ...}
// %v to print default format
// range: iterate over an array, it returns the index and the value

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// slice can have any size, can be sliced slice[low:high]

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	// slices cannot use equality operators
	// make: create a slice with a starting capacity of the length of a slice/array
	// sums := make([]int, 4)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
