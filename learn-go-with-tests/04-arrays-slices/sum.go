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

// slice can have any size
