package main

import (
	"fmt"
	"sort"
)

func combine(arr []float64, left float64) []float64 {
	sort.Float64s(arr)

	l := len(arr)

	comb := make([]float64, l/2)

	for i := 0; i < l/2; i++ {
		comb[i] = (arr[i] + arr[l-i-1]) * left
	}

	if l%2 != 0 {
		comb = append(comb, arr[l/2+1])
	}

	return comb
}

func main() {
	var n, p int

	fmt.Scan(&n)
	fmt.Scan(&p)

	left := 1 - float64(p)/100

	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for len(arr) > 1 {
		arr = combine(arr, left)
	}

	fmt.Println(arr[0])
}
