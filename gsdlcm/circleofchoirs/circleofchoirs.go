package main

import "fmt"

func gcd(a, b int) int {
	if b > 0 {
		return gcd(b, a%b)
	}

	return a
}

func main() {
	var n, m int

	for {
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			return
		}

		if gcd(n, m) == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
