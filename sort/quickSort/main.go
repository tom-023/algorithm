package main

import (
	"fmt"

	"github.com/alok87/goutils/pkg/random"
)

func main() {
	list := random.RangeInt(2, 100, 10)
	fmt.Println(list)
	fmt.Println(quickSort(list, 0, len(list)-1))
}

func quickSort(list []int, low int, high int) []int {
	if low < high {
		partitionIndex := partition(list, low, high)
		quickSort(list, low, partitionIndex-1)
		quickSort(list, partitionIndex+1, high)
	}

	return list
}

func partition(list []int, low int, high int) int {
	i := low - 1
	pivot := list[high]
	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i += 1
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[high] = list[high], list[i+1]

	return i + 1
}
