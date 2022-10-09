package main

import (
	"fmt"

	"github.com/alok87/goutils/pkg/random"
)

func main() {
	// list := []int{6, 34, 12, 2, 35, 3, 8, 4, 15, 31, 60, 54}
	list := random.RangeInt(2, 100, 5)
	bubbleSort(list)
	fmt.Printf("%v\n", list)
}

func bubbleSort(list []int) {
	n := len(list) - 1
	for j := 0; j < n; j++ {
		fmt.Printf("%v\n", list)
		for i := 0; i+j < n; i++ {
			fmt.Printf("%vラウンド, %v番目\n", j+1, i)
			left := list[i]
			right := list[i+1]
			if left > right {
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
	}
}
