package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n)

	ladder := make([]int, n+2)

	ladder[0] = 0
	ladder[n+1] = 0

	for i := 1; i <= n; i++ {
		fmt.Scan(&ladder[i])
	}

	fmt.Scan(&k)

	var localMax, j int

	for i := 1; i <= n+1; i++ {
		localMax = ladder[j]

		for j = i - k; j < i; j++ {
			if j < 0 {
				j = 0
			}

			if localMax < ladder[j] {
				localMax = ladder[j]
			}
		}

		ladder[i] = localMax + ladder[i]
	}

	fmt.Println(ladder[n+1])
}
