package main

import "fmt"

/*
1. Выведите перечень деликатесов — продуктов дороже 500 рублей.
2. Заказ выдан слайсом []string{"хлеб", "буженина", "сыр", "огурцы"}. Посчитайте стоимость заказа.
*/

func calcSum(prices map[string]float64, order []string) float64 {
	var total float64 = 0
	for _, product := range order {
		total += prices[product]
	}

	return total
}

func main() {
	prices := make(map[string]float64)
	prices["хлеб"] = 50
	prices["молоко"] = 100
	prices["масло"] = 200
	prices["колбаса"] = 500
	prices["огурцы"] = 200
	prices["сыр"] = 600
	prices["ветчина"] = 700
	prices["буженина"] = 900
	prices["помидоры"] = 250
	prices["рыба"] = 300
	prices["хамон"] = 1500

	// Выводим товары, цена которых выше 500 рублей
	for product, price := range prices {
		if price > 500 {
			fmt.Println(product)
		}
	}

	// Подсчитываем сумму заказа
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	fmt.Printf("Цена заказа: %d", int(calcSum(prices, order)))
}
