package main

import (
	"fmt"
	"sort"
	"strconv"
)

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func main() {
	list := make([]int, 3)

	for i := 0; ; i++ {
		var x string
		fmt.Printf("Enter an integer: ")
		_, _ = fmt.Scan(&x)
		if x == "X" {
			break
		}

		n, _ := strconv.Atoi(x)

		if i < 3 {
			pos := sort.SearchInts(list, 0)
			list[pos] = n
		} else {
			list = append(list, n)
		}

		sort.Ints(list)
		fmt.Println(list)
	}
}
