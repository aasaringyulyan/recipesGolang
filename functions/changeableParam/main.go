package main

import "fmt"

func main() {
	print5("1234567")
}

func print5(data string) {
	if len(data) > 5 {
		data = data[0:5]
	}

	fmt.Println(data)
}
