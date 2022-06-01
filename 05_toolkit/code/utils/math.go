package utils

import "fmt"

func printNumber(num int) {
	fmt.Println("Current Number: ", num)
}

// Add adds together multiple numbers.
func Add(nums ...int) int {
	total := 0
	for _, num := range nums {
		printNumber(num)
		total += num
	}
	return total
}
