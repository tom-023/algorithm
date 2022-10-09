package main

import "fmt"

func main() {
	list := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}
	fmt.Println(list)
	fmt.Println(binarySearch(list, 5))
}

func binarySearch(list []int, searchInt int) int {
	leftIndex, rightIndex := 0, len(list)-1
	for leftIndex <= rightIndex {
		fmt.Print("search\n")
		mid := (leftIndex + rightIndex) / 2
		if list[mid] == searchInt {
			return mid
		} else if list[mid] < searchInt {
			leftIndex = mid + 1
		} else {
			rightIndex = mid - 1
		}
	}

	return -1
}
