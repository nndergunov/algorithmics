package main

import (
	"fmt"
	"sort"
)

type time struct {
	h int
	m int
	s int
}

func compare(f, s time) bool {
	if f.h > s.h {
		return true
	}

	if f.h == s.h {
		if f.m > s.m {
			return true
		}

		if f.m == s.m {
			if f.s > s.s {
				return true
			}
		}
	}

	return false
}

func main() {
	var n int

	fmt.Scanln(&n)

	t := make([]time, n)

	for i := 0; i < n; i++ {
		t[i] = time{}

		fmt.Scan(&t[i].h)
		fmt.Scan(&t[i].m)
		fmt.Scan(&t[i].s)
	}

	sort.Slice(t, func(f, s int) bool { return !compare(t[f], t[s]) })

	for i := 0; i < n; i++ {
		fmt.Println(t[i].h, t[i].m, t[i].s)
	}
}
