package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func main() {
	var (
		n, m int
		s    [21][21]int
	)

	for {
		_, err := fmt.Scan(&n, &m)
		if err != nil {
			return
		}

		s[0][0] = 1

		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				s[i][j] = s[i-1][j-1] + (i-1)*s[i-1][j]
			}
		}

		a := 0

		for i := 1; i <= m; i++ {
			a += s[n][i]
		}

		b := 1

		for i := 1; i <= n; i++ {
			b *= i
		}

		d := gcd(a, b)
		a /= d
		b /= d

		fmt.Printf("%d/%d\n", a, b)
	}
}
