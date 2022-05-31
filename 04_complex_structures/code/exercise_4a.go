package main

import "fmt"

// func average(x, y, z float64) float64 {
// 	total := x + y + z
// 	return total / 3
// }

func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Println(average(10, 5, 7))
}
