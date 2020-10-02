package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func partialSort(c chan []int) {
	defer wg.Done()
	a := <-c
	fmt.Println(a)
	sort.Ints(a)
	c <- a
}

func mergeSorted(arr1, arr2 []int, c chan []int) {
	i, j, k := 0, 0, 0
	tempArr := make([]int, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			tempArr[k] = arr1[i]
			i++
		} else {
			tempArr[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		tempArr[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		tempArr[k] = arr2[j]
		j++
		k++
	}
	c <- tempArr
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]int, 0)
	fmt.Printf("Enter sequence of integers by space:\n")

	scanner.Scan()
	nums := scanner.Text()

	strNums := strings.Split(nums, " ")

	for i := range strNums {
		n, _ := strconv.Atoi(strNums[i])
		data = append(data, n)
	}

	length := len(data)
	partialLength := int(math.Ceil(float64(length) / 4.0))
	// fmt.Println("arr len: ", length, "part len: ", partialLength)

	c := make(chan []int, 4)

	wg.Add(4)
	c <- data[:partialLength]
	c <- data[partialLength : 2*partialLength]
	c <- data[2*partialLength : 3*partialLength]
	c <- data[3*partialLength:]

	go partialSort(c)
	go partialSort(c)
	go partialSort(c)
	go partialSort(c)

	wg.Wait()

	ch := make(chan []int)
	go mergeSorted(<-c, <-c, ch)
	go mergeSorted(<-c, <-c, ch)
	go mergeSorted(<-ch, <-ch, ch)

	fmt.Println(<-ch)
}
