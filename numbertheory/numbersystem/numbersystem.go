package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		n, m int
		a    string
	)

	fmt.Scan(&n, &m, &a)

	num, _ := strconv.ParseInt(a, n, 64)

	fmt.Println(strings.ToUpper(strconv.FormatInt(num, m)))
}
