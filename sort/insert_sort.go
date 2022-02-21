package main

// 插入排序：前面部分的数组是保持有序的，找到遍历元素在前面数组中的位置

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		// 当前遍历的元素
		k := i
		for ; k > 0 && arr[k-1] > tmp; k-- {
			arr[k] = arr[k-1]
		}
		arr[k] = tmp
	}
}
