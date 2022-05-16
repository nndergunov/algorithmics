package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)

	var max int

	for i := 0; i < n; i++ {
		if max < arr[i]-1 {
			break
		}

		max += arr[i]
	}

	max++

	fmt.Println(max)
}
