package main

import "fmt"

func search(x, y, n, c int, ch [][]rune) int {
	if ch[x][y] == '*' || x+1 == n || y+1 == n || x-1 == -1 || y-1 == -1 {
		return c
	}

	c++

	ch[x][y] = '*'

	c = search(x, y-1, n, c, ch)
	c = search(x, y+1, n, c, ch)
	c = search(x-1, y, n, c, ch)
	c = search(x+1, y, n, c, ch)

	return c
}

func main() {
	var x, y, n int

	fmt.Scan(&n)

	text := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&text[i])
	}

	ch := make([][]rune, n)

	for i := 0; i < n; i++ {
		ch[i] = make([]rune, len(text[i]))

		for j := 0; j < n; j++ {
			ch[i][j] = rune(text[i][j])
		}
	}

	fmt.Scan(&x, &y)
	c := search(x-1, y-1, n, 0, ch)

	fmt.Println(c)
}
