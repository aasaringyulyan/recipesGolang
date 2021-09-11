package main

import (
	"fmt"
)

func main() {
	var K, N int
	fmt.Scan(&K, &N)

	A := make([]int, 0, N)

	var n int

	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		A = append(A, n)
	}

	dist := A[len(A)-1] - A[0]

	for i := 0; i < len(A)-1; i++ {
		rez := 20 - (A[i+1] - A[i])

		if rez < dist {
			dist = rez
		}
	}

	fmt.Println(dist)
}
