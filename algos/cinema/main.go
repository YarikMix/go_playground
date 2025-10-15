package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	if (N == 0) || (M == 0) || (N > 2*M) || (M > 2*N) {
		fmt.Println("NO SOLUTION")
		return
	}

	answer := ""
	if N >= M {
		k := N - M
		for range k {
			answer += "BGB"
		}
		for range M - k {
			answer += "BG"
		}
	} else {
		k := M - N
		for range k {
			answer += "GBG"
		}
		for range N - k {
			answer += "GB"
		}
	}

	fmt.Println(answer)
}
