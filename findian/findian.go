package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Printf("Enter a string: ")
	_, _ = fmt.Scan(&str)

	strSmall := strings.ToLower(str)

	if strings.HasPrefix(strSmall, "i") &&
		strings.Contains(strSmall, "a") &&
		strings.HasSuffix(strSmall, "n") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
