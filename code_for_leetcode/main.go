package main

import (
	"fmt"
)

//func Solution(N int) {
//	var enable_print int;
//	enable_print = N % 10;
//	for N > 0 {
//		if enable_print == 0 && N % 10 != 0 {
//			enable_print = 1;
//		}
//		if enable_print != 0 {
//			fmt.Print(N % 10);
//		}
//		N = N / 10;
//	}
//}

func Solution(A []int, B []int, N int) int {
	cities := make([]int, N+1)

	for i := 0; i < len(A); i++ {
		cities[A[i]]++
		cities[B[i]]++
	}

	maxRank := 0

	for i := 0; i < len(A); i++ {
		rank := cities[A[i]] + cities[B[i]] - 1
		if rank > maxRank {
			maxRank = rank
		}
	}

	return maxRank
}


func main() {
	A := []int{1, 2, 3, 3}
	B := []int{2, 3, 1, 4}
	N := 4
	result := Solution(A, B, N)
	fmt.Println(result) // 应输出 4

	A = []int{1, 2, 4,5}
	B = []int{2, 3, 5, 6}
	N = 6
	result = Solution(A, B, N)
	fmt.Println(result) // 应输出 4

}
