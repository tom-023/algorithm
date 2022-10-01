package main

import "fmt"

func main() {
	list := []int{6, 34, 12, 2, 35, 3, 8, 4, 15, 31}

	buildHeap(list...)

	n := len(list) - 1
	for i := 0; i < n; i++ {
		heapify(0, n-i, list...)
		// 先頭と末尾の要素を入れ替える
		list[0], list[n-i] = list[n-i], list[0]
		fmt.Printf("---------------------------\n")
		fmt.Printf("list: %v\n", list)
		fmt.Printf("---------------------------\n")
	}

	fmt.Println(list)
}

func buildHeap(list ...int) {
	n := len(list) / 2 // ヒープの深さ
	for i := 0; i < n; i++ {
		heapify(i, len(list), list...)
	}
	fmt.Printf("buildHeap---------------------------\n")
	fmt.Printf("list: %v\n", list)
	fmt.Printf("buildHeap---------------------------\n")
}

func heapify(parentNum int, listSize int, list ...int) {
	/*
		list(parentNum)が親
		list(childLeft)が左の子
		list(childRight)が右の子
	*/
	largest := parentNum
	childLeft := parentNum*2 + 1
	childRight := childLeft + 1

	// 親と左子を比較し、子の方が大きかったら親と子を入れ替える
	if childLeft < listSize && list[childLeft] > list[largest] {
		largest = childLeft
	}

	if childRight < listSize && list[childRight] > list[largest] {
		largest = childRight
	}

	if largest != parentNum {
		// 親と子の値を入れ替える
		list[parentNum], list[largest] = list[largest], list[parentNum]
		heapify(0, listSize, list...)
	}
}
