package main

import (
	"fmt"

	"github.com/alok87/goutils/pkg/random"
)

func main() {
	list := random.RangeInt(2, 100, 10)
	fmt.Println(selectSort(list))
}

func selectSort(list []int) []int {
	n := len(list)

	for i := range list {
		minIndex := i
		fmt.Printf("%vラウンド\n", i+1)
		fmt.Printf("list: %v\n", list)
		fmt.Printf("minIndex: %v\n", minIndex)

		for j := 0; j+i < n; j++ {
			fmt.Printf("%v番目\n", j+i)
			if list[minIndex] > list[j+i] {
				minIndex = j + i
			}
		}

		list[i], list[minIndex] = list[minIndex], list[i]
	}

	return list
}
