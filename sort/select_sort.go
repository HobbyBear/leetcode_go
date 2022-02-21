package main

import "fmt"

// 选择排序：遍历整个数组，每次选出最小的元素与元素前面的位置进行交换。

func main() {
	arr := []int{7, 5, 3, 3, 10}
	BubbleSort(arr)
	fmt.Println(arr)
}

func SelectSort(arr []int) {
	if len(arr) == 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		swap(arr, minIndex, i)
	}
}

func swap(arr []int, i, j int) {
	num := arr[i]
	arr[i] = arr[j]
	arr[j] = num
}
