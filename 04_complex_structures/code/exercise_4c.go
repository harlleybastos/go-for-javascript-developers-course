package main

import (
	"fmt"
)

func calculateAverage(scores []float64) float64 {
	total := 0.0
	for _, value := range scores {
		total += value
	}
	return total / float64(len(scores))
}

var initialPets map[string]string = map[string]string{
	"fido":     "dog",
	"penelope": "cat",
	"neo":      "horse",
}

var initialGroceries = []string{
	"milk",
	"eggs",
	"bread",
	"apples",
}

func addGroceryToList(newGroceries ...string) []string {
	foods := initialGroceries
	for _, food := range newGroceries {
		foods = append(foods, food)
	}
	return foods
}

func doesPetExist(petName string) bool {
	_, exists := initialPets[petName]
	return exists
}

func main() {
	// scores := float64{97, 92, 75, 83, 91, 70}
	// fmt.Println(calculateAverage(scores))

	// pet := "fido"
	// petExist := doesPetExist(pet)
	// fmt.Println(petExist)

	groceries := addGroceryToList("bananas", "tomatoes")
	fmt.Println(groceries)
}
