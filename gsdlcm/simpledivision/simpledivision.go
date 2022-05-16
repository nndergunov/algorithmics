package main

import (
	"fmt"
	"math"
)

func gcd(a, b int) int {
	if b > 0 {
		return gcd(b, a%b)
	}

	return a
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func main() {
	var a, b int

	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			return
		}

		res := abs(a - b)
		a = b

		for {
			_, err = fmt.Scan(&b)
			if err != nil || b == 0 {
				break
			}

			res = gcd(res, abs(a-b))
			a = b
		}

		fmt.Printf("%d\n", res)
	}
}
