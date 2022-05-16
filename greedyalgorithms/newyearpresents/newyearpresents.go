package main

import (
	"fmt"
)

type present struct {
	packing    int
	delivering int
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func compare(a, b present) bool {
	return min(a.packing, a.delivering) > min(b.packing, b.delivering)
}

func sort(array []present) []present {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if compare(array[i], array[j]) {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

func main() {
	var n int

	fmt.Scan(&n)

	presents := make([]present, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&presents[i].packing)
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&presents[i].delivering)
	}

	presents = sort(presents)
	arr := make([]present, n)
	i := 0
	j := n - 1

	for k := 0; k < n; k++ {
		if presents[k].packing < presents[k].delivering {
			arr[i] = presents[k]
			i++
		} else {
			arr[j] = presents[k]
			j--
		}
	}

	sum, firstSum := 0, 0

	for i := 0; i < n; i++ {
		firstSum += arr[i].packing
		sum = max(sum, firstSum) + arr[i].delivering
	}

	fmt.Println(sum)
}
