package main

import "fmt"

func main() {
	var sum = 0

	getSum(&sum, 5, 3)

	fmt.Println("sum =", sum)
}

func getSum(sum *int, n1 int, n2 int) {
	*sum = n1 + n2
}
