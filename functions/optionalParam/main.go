package main

import "fmt"

func main() {
	sayGoodby1("")
	sayGoodby1("see you")

	sayGoodby2()
	sayGoodby2("see you")
}

func sayGoodby2(message ...string) {
	if len(message) == 0 {
		fmt.Println("Goodby!")
	} else {
		fmt.Println(message[0])
	}
}

func sayGoodby1(message string) {
	if len(message) == 0 {
		fmt.Println("Goodby!")
	} else {
		fmt.Println(message)
	}
}
