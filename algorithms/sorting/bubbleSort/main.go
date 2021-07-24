package main

import (
	"fmt"
	"time"
)

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := bubbleSort(items)

	// *** simplified speed test ***

	items = make([]int, 200)

	for i := 0; i < len(items); i++ {
		items[i] = i
	}

	items[5], items[6] = items[6], items[5]
	count := 10000
	start := time.Now()

	for i := 0; i < count; i++ {
		bubbleSort(items)
	}

	delta := time.Now().Sub(start)
	nanosecond := delta.Nanoseconds()

	fmt.Println(sortItems)
	fmt.Println(nanosecond)
}

func bubbleSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)

	copy(items, arr)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if items[j] < items[i] {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	return items
}
