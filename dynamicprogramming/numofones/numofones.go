package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func counter(n int, c map[int]int) int {
	if n == 1 {
		return 1
	}

	if _, ok := c[n]; ok {
		return c[n]
	}

	m := 1 + counter(n-1, c)

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			m = min(m, counter(i, c)+counter(n/i, c))
		}
	}

	c[n] = m

	return m
}

func main() {
	var n int

	fmt.Scan(&n)

	c := make(map[int]int)

	fmt.Println(counter(n, c))
}
