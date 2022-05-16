package main

import "fmt"

func count(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	res := make([]int, n+1)

	res[0] = 1
	res[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			res[i] += res[j] * res[i-j-1]
		}
	}

	return res[n]
}

func main() {
	var n int

	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}

		fmt.Println(count(n / 2))
	}
}
