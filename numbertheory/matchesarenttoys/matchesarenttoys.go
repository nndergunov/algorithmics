package main

import "fmt"

func main() {
	var t int

	fmt.Scan(&t)

	var ans string

	for i := 0; i < t; i++ {
		var n, m int

		fmt.Scan(&n)
		fmt.Scan(&m)

		k := n - (n/(1+m))*(1+m)
		if k <= m && k != 0 {
			ans += "1"
		} else {
			ans += "2"
		}
	}

	fmt.Println(ans)
}
