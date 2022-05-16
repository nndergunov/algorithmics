package main

import (
	"fmt"
	"sort"
	"strconv"
)

type student struct {
	surname   string
	name      string
	class     string
	birthdate string
}

func compClass(f, s string) int {
	fNum, _ := strconv.Atoi(f[:len(f)-1])
	sNum, _ := strconv.Atoi(s[:len(s)-1])
	fLett := f[len(f)-1]
	sLett := s[len(s)-1]

	if fNum < sNum {
		return 1
	}

	if fNum == sNum {
		if fLett < sLett {
			return 1
		}

		if fLett == sLett {
			return 0
		}
	}

	return -1
}

func less(f, s student) bool {
	if compClass(f.class, s.class) == 1 {
		return true
	}

	if compClass(f.class, s.class) == 0 {
		if f.surname < s.surname {
			return true
		}
	}

	return false
}

func main() {
	var n int

	fmt.Scan(&n)

	s := make([]student, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i].surname)
		fmt.Scan(&s[i].name)
		fmt.Scan(&s[i].class)
		fmt.Scan(&s[i].birthdate)
	}

	sort.Slice(s, func(i, j int) bool { return less(s[i], s[j]) })

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %s\n", s[i].class, s[i].surname, s[i].name, s[i].birthdate)
	}
}
