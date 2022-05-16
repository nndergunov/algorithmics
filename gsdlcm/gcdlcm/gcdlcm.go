package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)

	ans := ""

	for i := 0; i < n; i++ {
		var a, b int

		fmt.Scan(&a, &b)

		if b%a == 0 {
			ans = ans + strconv.Itoa(a) + " " + strconv.Itoa(b)
		} else {
			ans += "-1"
		}

		if i < n-1 {
			ans = ans + "\n"
		}
	}

	fmt.Println(ans)
}
