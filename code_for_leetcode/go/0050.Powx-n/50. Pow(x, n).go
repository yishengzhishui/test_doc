package leetcode

// 递归
func myPow(x float64, n int) float64 {
	// 递归终止条件
	if n == 0 {
		return 1
	}

	// 处理负指数
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	// 处理奇数指数
	if n%2 == 1 {
		return x * myPow(x, n-1)
	}

	// 处理偶数指数
	return myPow(x*x, n/2)
}

// 迭代
// https://leetcode.cn/problems/powx-n/solutions/241471/50-powx-n-kuai-su-mi-qing-xi-tu-jie-by-jyd/
// 其实是分治的思想
// x^n=(x^2)^(n/2)
// n/2 需要区分奇数和偶数
// n为偶数，(x^2)^(n//2)
// n为奇数时  n/2是向下取整的，所以要加一位 x*(x^2)^(n//2)
func myPow1(x float64, n int) float64 {
	// 如果指数 n 为负数，取绝对值，并将底数 x 取倒数
	if n < 0 {
		n = -n
		x = 1 / x
	}

	result := 1.0 // 初始化结果变量为 1

	// 迭代计算幂次
	for n > 0 {
		// 如果当前二进制位为 1，累乘当前底数 x 到结果中
		if n%2 == 1 {
			result *= x
		}

		// 底数 x 自身平方，相当于将指数 n 减半
		x *= x

		// 右移一位，相当于整除 2
		n = n >> 1
	}

	return result
}
