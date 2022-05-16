package main

import "fmt"

func farrey(n, m, a1, b1, a2, b2 int) {
	if n < b1+b2 {
		return
	}

	if m == a1+a2 && n == b1+b2 {
		return
	}

	if m*(b1+b2) < n*(a1+a2) {
		fmt.Print("R")
		farrey(n, m, a1, b1, a1+a2, b1+b2)
	} else {
		fmt.Print("L")
		farrey(n, m, a1+a2, b1+b2, a2, b2)
	}
}

func main() {
	var n, m int

	for {
		_, err := fmt.Scan(&n, &m)
		if err != nil || (n == 1 && m == 1) {
			return
		}

		farrey(n, m, 0, 1, 1, 0)
		fmt.Print("\n")
	}
}
