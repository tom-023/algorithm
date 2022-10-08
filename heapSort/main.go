package main

import (
	"fmt"

	"github.com/alok87/goutils/pkg/random"
)

func main() {
	// list := []int{6, 34, 12, 2, 35, 3, 8, 4, 15, 31, 60, 54}
	list := random.RangeInt(2, 100, 10)
	// ヒープを構築
	buildHeap(list)

	fmt.Printf("初期のヒープ構築---------------------------\n")
	fmt.Printf("%v\n", list)
	fmt.Printf("初期のヒープ構築---------------------------\n")

	n := len(list) - 1
	for i := 0; i < n; i++ {
		// 最も大きい数値の先頭と末尾の要素を入れ替える（根を取り出し、一番最後の要素を根にセットする）
		list[0], list[n-i] = list[n-i], list[0]

		fmt.Printf("Before Heap---------------------------\n")
		fmt.Printf("%v\n", list)
		fmt.Printf("Before Heap---------------------------\n")

		// ヒープを再構築
		heapify(0, n-i, list)

		fmt.Printf("After Heap---------------------------\n")
		fmt.Printf("%v\n", list)
		fmt.Printf("After Heap---------------------------\n")

	}

	fmt.Println(list)
}

func buildHeap(list []int) {
	n := len(list)/2 - 1 // 一番最後の親の添字

	// ヒープを構築する（一番最後の親から一番上の親（根）の順で検証する）
	for ; n >= 0; n-- {
		fmt.Printf("---------------------------\n")
		fmt.Printf("親 %v のソート\n", list[n])
		fmt.Printf("---------------------------\n")

		heapify(n, len(list), list)
	}
}

func heapify(parentNum int, listSize int, list []int) {
	/*
		list(parentNum)が親
		list(childLeft)が左の子
		list(childRight)が右の子
	*/
	largest := parentNum
	childLeft := parentNum*2 + 1
	childRight := parentNum*2 + 2

	// 親と左子を比較し、子の方が大きかったら親と子を入れ替える
	if (childLeft < listSize) && (list[childLeft] > list[largest]) {
		largest = childLeft
	}

	// （親 or 左子）と右子を比較し、右子の方が大きかったらlargestに右子の添字を設定する
	if (childRight < listSize) && (list[childRight] > list[largest]) {
		largest = childRight
	}

	if largest != parentNum {
		// 親より子の方が大きいため、親と子の値を入れ替える
		list[parentNum], list[largest] = list[largest], list[parentNum]

		heapify(largest, listSize, list)
	}
}
