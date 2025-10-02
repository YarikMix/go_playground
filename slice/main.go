package main

import "fmt"

func main() {
	dim := 100
	slice := make([]int, 0, dim)

	// заполняем slice числами
	for i := 0; i < 100; i++ {
		slice = append(slice, i+1)
	}

	// оставляем первые и последние 10 элементов
	slice = append(slice[:10], slice[dim-10:]...)

	dim = len(slice)

	// разворачиваем слайс
	for i := range slice[:dim/2] {
		slice[i], slice[dim-i-1] = slice[dim-i-1], slice[i]
	}

	fmt.Println(slice)
}
