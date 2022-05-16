package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		notes          = []int{10, 20, 50, 100, 200, 500}
		req, sum, took int
	)

	sort.Sort(sort.Reverse(sort.IntSlice(notes)))

	fmt.Scan(&req)

	for i := 0; i < len(notes); {
		if sum+notes[i] > req {
			i++

			continue
		}

		took++

		sum += notes[i]

		if sum == req {
			fmt.Println(took)

			return
		}
	}

	fmt.Println(-1)
}
