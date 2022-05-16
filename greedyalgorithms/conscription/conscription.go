package main

import "fmt"

type orc struct {
	g, h int
}

func compare(first, second *orc) bool {
	return (first.g - first.h) > (second.g - second.h)
}

func sort(arr []*orc) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if compare(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	var n, g, h int

	fmt.Scan(&n)
	fmt.Scan(&g)
	fmt.Scan(&h)

	if n < g+h {
		fmt.Println(-1)

		return
	}

	candidates := make([]*orc, n)

	for i := 0; i < n; i++ {
		var currG, currH int

		fmt.Scan(&currG)
		fmt.Scan(&currH)

		candidates[i] = &orc{
			g: currG,
			h: currH,
		}
	}

	sort(candidates)

	res := 0

	for i := 0; i < n; i++ {
		if i < h {
			res += candidates[i].h
		} else if i < n-g {
			if candidates[i].h > candidates[i].g {
				res += candidates[i].h
			} else {
				res += candidates[i].g
			}
		} else {
			res += candidates[i].g
		}
	}

	fmt.Println(res)
}
