package main

import "fmt"

func main() {
	var x, n int

	fmt.Scan(&n)

	m := make([]int, n)
	out := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)

		m[i] = x
	}

	for i := 0; i < n; i++ {
		out[m[m[i]-1]-1] = m[i]
	}

	for i := 0; i < n; i++ {
		fmt.Print(out[i], " ")
	}

	fmt.Print("\n")
}
