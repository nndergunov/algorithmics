package main

import "fmt"

func get(x, n int, b []int) (int, []int) {
	for i := n - 1; i >= 0; i-- {
		x -= b[i]
		if x == 0 {
			b[i] = 0

			return i + 1, b
		}
	}

	return 0, b
}

func main() {
	var n int

	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)
	f := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&f[i])

		b[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		a[i], b = get(f[i]+1, n, b)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}

	fmt.Printf("\n")
}
