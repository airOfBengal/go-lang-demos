package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Sorts an integer slice
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

// Swaps the current value in a slice with its +1 neighbour
func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nums := make([]int, 0, 10)

	// Get input from user and put the data in an integer slice
	fmt.Printf("Type up to ten integers (space separated): ")
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	for i := 0; i < len(input); i++ {
		n, _ := strconv.Atoi(input[i])
		nums = append(nums, n)
	}

	// Sort the slice
	BubbleSort(nums)

	// Print the sorted slice
	fmt.Println(nums)
}