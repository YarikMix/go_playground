package main

import "fmt"

/*
Дан массив целых чисел A и целое число k. Нужно найти и вывести индексы пары чисел, сумма которых равна k. Если таких чисел нет, то вернуть пустой слайс. Индексы можно вернуть в любом порядке.
*/

func solve(A []int, k int) []int {
	// Создаем мапу для хранения чисел и их индексов
	numMap := make(map[int]int)

	for i, num := range A {
		// Вычисляем число, которое нужно найти для текущего числа
		complement := k - num

		// Проверяем, есть ли complement в мапе
		if idx, exists := numMap[complement]; exists {
			// Если нашли пару, возвращаем индексы
			return []int{i, idx}
		}

		// Сохраняем текущее число и его индекс в мапу
		numMap[num] = i
	}

	return nil
}

func main() {
	arr := []int{2, 5, 6, 7, 8}

	fmt.Print(solve(arr, 11))
}
