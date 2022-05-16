package main

import "fmt"

func gcd(a, b int) int {
	if b > 0 {
		return gcd(b, a%b)
	}

	return a
}

func main() {
	for {
		var n int

		fmt.Scan(&n)

		if n == 0 {
			break
		}

		g := 0

		for i := 1; i < n; i++ {
			for j := i + 1; j <= n; j++ {
				g += gcd(i, j)
			}
		}

		fmt.Println(g)
	}
}
