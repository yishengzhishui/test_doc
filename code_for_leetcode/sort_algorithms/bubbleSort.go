package main

import "fmt"

//嵌套循环，每次查看相邻的元素，如果逆序则交换。
func bubbleSort(arr []int) []int {
	n := len(arr)            // 获取切片的长度
	for i := 0; i < n; i++ { // 外层循环，控制每次循环中最大元素的位置
		for j := 0; j < n-i-1; j++ { // 内层循环，相邻元素的比较和交换
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // 如果相邻元素逆序，则交换它们的位置
			}
		}
	}
	return arr // 返回排序后的数组
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("原始数组:", arr)

	bubbleSort(arr)

	fmt.Println("冒泡排序后数组:", arr)
}
