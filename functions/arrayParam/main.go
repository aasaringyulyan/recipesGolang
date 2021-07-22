package main

import "fmt"

func main() {
	var avg = getAvg(1, 2, 3, 4)

	fmt.Println("avg =", avg)
}

func getAvg(values ...float64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum = 0.0

	for _, value := range values {
		sum += value
	}

	return sum / float64(len(values))
}
