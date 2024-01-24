package main

import "fmt"

//从前到后逐步构建有序序列，对于未排序的数据，在已排序序列中从后向前扫描，找到相应位置并插入
func insertionSort(arr []int) []int {
	n := len(arr)  // 获取切片的长度
	for i := 1; i < n; i++ {  // 外层循环，从第二个元素开始遍历整个数组
		value := arr[i]  // 将当前元素保存到 value 中
		j := i - 1  // j 为已排序部分的最后一个元素的索引
		for j >= 0 && value < arr[j] {  // 内层循环，在已排序的部分中找到合适的插入位置
			arr[j+1] = arr[j]  // 将大于 value 的元素后移一位
			j--
		}
		arr[j+1] = value  // 插入 value 到正确的位置
	}
	return arr  // 返回排序后的数组
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("原始数组:", arr)

	insertionSort(arr)

	fmt.Println("插入排序后数组:", arr)
}
