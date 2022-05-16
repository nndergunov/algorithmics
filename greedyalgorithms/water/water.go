package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, res int

	fmt.Scan(&n, &k)

	canisters := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&canisters[i])

		if canisters[i] > k {
			fmt.Println("Impossible")

			return
		}
	}

	sort.Ints(canisters)

	i, j := 0, 1

	for i+j < n {
		res++

		if canisters[i]+canisters[n-j] <= k {
			i++
		}

		j++
	}

	if i+j != n {
		res++
	}

	fmt.Println(res)
}
