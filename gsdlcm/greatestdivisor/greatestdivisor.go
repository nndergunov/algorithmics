package main

import "fmt"

func gcd(a, b int) int {
	if b > 0 {
		return gcd(b, a%b)
	}

	return a
}

func main() {
	var n int

	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	g := arr[0]

	for i := 1; i < n; i++ {
		g = gcd(g, arr[i])
	}

	fmt.Println(g)
}
