package main

import "fmt"

func main() {
	sentenceValue := "Sentence value"

	for _, value := range sentenceValue {

		fmt.Println(string(value))
	}

}
