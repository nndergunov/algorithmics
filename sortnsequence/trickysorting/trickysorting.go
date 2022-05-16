package main

import (
	"fmt"
	"sort"
)

func less(f, s int) bool {
	if f%10 < s%10 {
		return true
	}

	if f%10 == s%10 {
		if f < s {
			return true
		}
	}

	return false
}

func main() {
	var n int

	fmt.Scan(&n)

	num := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}

	sort.Slice(num, func(f, s int) bool { return less(num[f], num[s]) })

	for i := 0; i < n; i++ {
		fmt.Print(num[i], " ")
	}
}
