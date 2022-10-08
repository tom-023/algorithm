package main

import (
	"fmt"

	"github.com/alok87/goutils/pkg/random"
)

func main() {
	list := random.RangeInt(2, 100, 10)
	fmt.Println(list)
	fmt.Println(insertionSort(list...))
}

func insertionSort(list ...int) []int {
	n := len(list)

	for i := 1; i < n; i++ {
		tmp := list[i]

		j := i - 1
		for j >= 0 && list[j] > tmp {
			list[j+1] = list[j]
			j -= 1
		}

		list[j+1] = tmp
	}

	return list
}
