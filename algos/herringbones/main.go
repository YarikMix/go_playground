package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	square := n*n + 3*n + 1
	fmt.Print(square)
}
