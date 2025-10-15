package main

import "fmt"

func main() {
	var k, m, a, b int
	fmt.Scan(&k, &m, &a, &b)

	katya := b/k - (a-1)/k
	masha := b/m - (a-1)/m
	fmt.Print(katya - masha)
}
