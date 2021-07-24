package main

import (
	"fmt"
	"time"
)

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := quickSort(items)

	// *** simplified speed test  ***

	items = make([]int, 200)

	for i := 0; i < len(items); i++ {
		items[i] = i
	}

	items[5],items[6] = items[6], items[5]
	count := 10000
	start := time.Now()

	for i := 0; i < count; i++ {
		quickSort(items)
	}

	delta := time.Now().Sub(start)
	nanosecond := delta.Nanoseconds()

	fmt.Println(sortItems)
	fmt.Println(nanosecond)
}

func quickSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)

	copy(items, arr)
	doSort(items, 0, length-1)

	return items
}

func doSort(items []int, first int, last int) {
	if first > last || len(items) == 0 {
		return
	}

	i := first
	j := last
	x := items[(first+last)/2]

	for i <= j {
		for items[i] < x {
			i++
		}

		for items[j] > x {
			j--
		}

		if i <= j {
			items[i], items[j] = items[j], items[i]
			i++
			j--
		}
	}

	doSort(items, first, j)
	doSort(items, i, last)
}
