package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToMin(t string) int {
	parts := strings.Split(t, ":")

	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	return h*60 + m
}

type visit struct {
	from int
	to   int
}

func newVisit(from, to string) *visit {
	return &visit{
		from: convertToMin(from),
		to:   convertToMin(to),
	}
}

func sort(v []*visit) {
	for i := 0; i < len(v); i++ {
		for j := i + 1; j < len(v); j++ {
			if v[i].to > v[j].to {
				v[i], v[j] = v[j], v[i]
			}
		}
	}
}

func main() {
	var n int

	fmt.Scan(&n)

	vis := make([]*visit, n)

	for i := 0; i < n; i++ {
		var from, to string

		fmt.Scan(&from)
		fmt.Scan(&to)

		vis[i] = newVisit(from, to)
	}

	sort(vis)

	free := vis[0].to

	var res int

	for i := 1; i < n; i++ {
		if vis[i].from >= free {
			res++

			free = vis[i].to
		}
	}

	res++

	fmt.Println(res)
}
