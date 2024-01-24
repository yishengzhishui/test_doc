package main

import "fmt"

//每次找到最小（大）值，然后放到待排序数组的起始位置

func selectionSort(arr []int) []int {
	n := len(arr)            // 获取切片的长度
	for i := 0; i < n; i++ { // 外层循环，遍历整个数组
		minIndex := i            // 假设当前位置为最小值的索引
		for j := i; j < n; j++ { // 内层循环，从未排序的部分找到最小值的索引
			if arr[minIndex] > arr[j] {
				minIndex = j // 更新最小值的索引
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i] // 将当前位置的元素与最小值进行交换
	}
	return arr // 返回排序后的数组
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("原始数组:", arr)

	selectionSort(arr)

	fmt.Println("选择排序后数组:", arr)
}
