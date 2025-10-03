package main

import "fmt"

// Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса

func RemoveDuplicates(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]bool, len(input))

	for _, val := range input {
		if _, exists := inputSet[val]; !exists {
			output = append(output, val)
		}

		inputSet[val] = true
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
