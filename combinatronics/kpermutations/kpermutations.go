package main

import (
	"fmt"
	"math"
)

func count(p []int, k, d int, s map[int]interface{}) int {
	l := len(p)

	if d == l {
		return 1
	}

	cem := 0

	for i := 1; i <= l; i++ {
		if _, ok := s[i]; !ok && (d == 0 || math.Abs(float64(p[d-1]-i)) <= float64(k)) {
			p[d] = i
			s[i] = nil
			cem += count(p, k, d+1, s)

			delete(s, i)
		}
	}

	return cem
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	p := make([]int, n)
	s := make(map[int]interface{})

	fmt.Println(count(p, k, 0, s))
}
