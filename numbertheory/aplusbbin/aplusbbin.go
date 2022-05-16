package main

import "fmt"

func max(f, s int) int {
	if f > s {
		return f
	}

	return s
}

func sum(a, b []byte) string {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	carry := 0

	for i := 0; i < max(len(a), len(b)) || carry != 0; i++ {
		if i == len(a) {
			a = append(a, byte('0'))
		}

		var add byte

		if i < len(b) {
			add = b[i] - byte(48)
		} else {
			add = 0
		}

		a[i] += byte(carry) + add

		if a[i] > '1' {
			carry = 1
		} else {
			carry = 0
		}

		if carry != 0 {
			a[i] -= 2
		}
	}

	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}

	return string(a)
}

func main() {
	var n, m string

	fmt.Scan(&n, &m)

	fmt.Println(sum([]byte(n), []byte(m)))
}
