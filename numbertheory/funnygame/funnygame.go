package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)

	max := n
	fmtd := strconv.FormatInt(int64(n), 2)

	l := len(fmtd)

	for i := 0; i < l; i++ {
		fmtd = string(fmtd[l-1]) + fmtd[:l-1]
		curr, _ := strconv.ParseInt(fmtd, 2, 64)

		if int(curr) > max {
			max = int(curr)
		}
	}

	fmt.Println(max)
}
