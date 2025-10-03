package main

import (
	"fmt"
	"slices"
)

// Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса

func RemoveDuplicates(input []string) []string {
	output := make([]string, 0)

	for _, val := range input {
		if !slices.Contains(output, val) {
			output = append(output, val)
		}
	}

	return output
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println(RemoveDuplicates(input))
}
