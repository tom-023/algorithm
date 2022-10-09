package main

import (
	"fmt"

	"github.com/alok87/goutils/pkg/random"
)

func main() {
	list := random.RangeInt(2, 100, 8)
	// list := []int{5, 4, 1, 8, 7, 3, 2, 9}
	fmt.Println(list)
	fmt.Println(mergeSort(list))
}

func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	half := len(list) / 2

	leftOrigin := list[:half]
	leftLen := len(leftOrigin)
	// listが変更されるとleftも変更する（leftはlistを参照している）ため、新しいリスト型のスライスを作成する
	left := make([]int, leftLen)
	copy(left, leftOrigin)

	rightOrigin := list[half:]
	rightLen := len(rightOrigin)
	right := make([]int, rightLen)
	copy(right, rightOrigin)

	mergeSort(left)
	mergeSort(right)

	var i, j, k int // i: leftのindexを管理する, j: rightのindexを管理する, k: leftとrightを結合したスライスのindexを管理する
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			list[k] = left[i]
			i += 1
		} else {
			list[k] = right[j]
			j += 1
		}
		k += 1
	}

	for i < len(left) {
		list[k] = left[i]
		i += 1
		k += 1
	}

	for j < len(right) {
		list[k] = right[j]
		j += 1
		k += 1
	}

	return list
}
