package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	romans := make([]int, n)

	for i := 0; i < n; i++ {
		romans[i] = i + 1
	}

	d := -1

	for len(romans) > 1 {
		for j := 0; j < k; j++ {
			d++

			if d == len(romans) {
				d = 0
			}
		}

		romans = append(romans[:d], romans[d+1:]...)

		d--
	}

	fmt.Println(romans[0])
}
