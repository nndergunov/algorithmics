package main

import (
	"fmt"
	"math"
)

func min(f, s int) int {
	if f < s {
		return f
	}

	return s
}

func main() {
	var (
		n, k int
		p    = 2
		mn   = math.MaxInt32
		v    []int
	)

	mp := make(map[int]int)
	mp1 := make(map[int]int)

	fmt.Scan(&n, &k)

	for p <= k {
		if k%p == 0 {
			h := 0

			for k%p == 0 {
				k /= p
				h++
			}

			mp1[p] = h
			v = append(v, p)
		}

		p++
	}

	for i := 0; i < len(v); i++ {
		num := n
		for num != 0 {
			mp[v[i]] += num / v[i]
			num /= v[i]
		}
	}
	
	for i := 0; i < len(v); i++ {
		mn = min(mn, mp[v[i]]/mp1[v[i]])
	}

	fmt.Println(mn)
}
