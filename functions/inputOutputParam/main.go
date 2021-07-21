package main

import "fmt"

func main() {
	var s1 = "A"
	var s2 = "B"

	swapStrings(&s1, &s2)

	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
}

func swapStrings(s1, s2 *string) {
	*s1, *s2 = *s2, *s1
}
