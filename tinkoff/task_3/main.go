package main

import "fmt"

func main() {
	var K, N int
	fmt.Scan(&N, &K)

	A := make([]int, 0, N)

	var n int

	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		A = append(A, n)
	}

	l := 1
	r := findMax(A)

	for l <= r {
		mid := l + (r-l)/2

		if isValid(A, N, mid, K) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	fmt.Println(r)
}

func isValid(a []int, n int, mid int, k int) bool {
	var count int

	for i := 0; i < n; i++ {
		count += a[i] / mid
	}

	return count >= k
}

func findMax(a []int) (max int) {
	max = a[0]

	for _, value := range a {
		if value > max {
			max = value
		}
	}

	return max
}
