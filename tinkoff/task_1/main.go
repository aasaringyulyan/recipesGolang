package main

import (
	"fmt"
	"strings"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var S string
		fmt.Scan(&S)

		str := strings.Split(S, "")

		if (str[0] == str[1] && str[2] == str[3]) ||
			(str[0] == str[2] && str[1] == str[3]) ||
			(str[0] == str[3] && str[1] == str[2]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
