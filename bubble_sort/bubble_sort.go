package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, position int) {
	temp := slice[0]
	slice[0] = slice[1]
	slice[1] = temp
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr[j:j+2], j)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]int, 0)
	fmt.Printf("Enter sequence of integers by space up to 10 numbers:\n")

	scanner.Scan()
	nums := scanner.Text()

	strNums := strings.Split(nums, " ")

	for i := range strNums {
		n, _ := strconv.Atoi(strNums[i])
		data = append(data, n)
	}

	fmt.Println("Original nums:\n", data)

	BubbleSort(data[:])
	fmt.Println("After Bubble Sort:\n", data)
}
