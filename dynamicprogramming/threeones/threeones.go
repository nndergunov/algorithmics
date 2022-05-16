package main

import "fmt"

func main() {
	var (
		n   int
		a1  = 2
		a2  = 4
		a3  = 7
		a4  = 0
		mod = 12345
	)

	fmt.Scan(&n)

	for i := 3; i < n; i++ {
		a4 = a1 + a2 + a3%mod
		a1 = a2 % mod
		a2 = a3 % mod
		a3 = a4 % mod
	}

	if n == 1 {
		fmt.Println(2)
	} else if n == 2 {
		fmt.Println(4)
	} else if n == 3 {
		fmt.Println(7)
	} else {
		fmt.Println(a4 % mod)
	}
}
