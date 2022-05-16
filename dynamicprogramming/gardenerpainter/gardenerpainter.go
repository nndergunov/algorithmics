package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	res := 3

	for i := 1; i < n; i++ {
		res *= 2
	}

	fmt.Println(res)
}
