package main

import "fmt"

func bin(a int) int {
	if a == 1 {
		return 2
	}

	if a == 2 {
		return 4
	}

	return bin(a-1) + bin(a-2)
}

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Println(bin(n))
}
