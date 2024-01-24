package main

import "fmt"

// 归并：先排序左右子数组，然后合并两个有序子数组
// 1. 把长度为n的输入序列分成两个长度为n/2的子序列；
//2. 对这两个子序列分别采用归并排序；
//3. 将两个排序好的子序列合并成一个最终的排序序列

// 归并排序的主函数
func mergeSort(arr []int, left, right int) {
	if right <= left {
		return
	}
	mid := (left + right) >> 1
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

// 归并排序中的合并操作
func merge(arr []int, left, mid, right int) {
	temp := make([]int, 0)
	i, j := left, mid+1

	// 比较两个有序数组的元素，将较小的元素添加到临时数组中
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	// 处理前一个数组未走完的情况
	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}

	// 处理后一个数组未走完的情况
	for j <= right {
		temp = append(temp, arr[j])
		j++
	}

	// 将临时数组中的有序元素复制回原数组
	copy(arr[left:right+1], temp)
}

// 主函数，入口点
func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("原始数组:", arr)

	// 调用归并排序函数对数组进行排序
	mergeSort(arr, 0, len(arr)-1)

	// 打印排序后的数组
	fmt.Println("归并排序后数组:", arr)
}
