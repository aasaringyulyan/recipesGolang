package main

import "fmt"

func main() {
	var sum = getSum(5, 3)

	fmt.Println(sum)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}
