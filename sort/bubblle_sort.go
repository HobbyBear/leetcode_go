package main

// 冒泡排序： 不断两两交换，将大的元素交换到末尾

func BubbleSort(arr []int) {
	j := len(arr) - 1
	for k := 0; k <= j; j-- {
		for i := 0; i < j; i++ {
			if arr[i+1] >= arr[i] {
				continue
			}
			tmp := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = tmp
		}
	}

}
