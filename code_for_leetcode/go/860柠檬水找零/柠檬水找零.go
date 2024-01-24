package leetcode

func lemonadeChange(bills []int) bool {
	// 初始化 5 元和 10 元的数量为 0
	five, ten := 0, 0

	// 遍历每个顾客支付的账单
	for _, bill := range bills {
		switch bill {
		case 5:
			// 如果是 5 元，直接增加 5 元的数量
			five++
		case 10:
			// 如果是 10 元，需要找零 5 元
			// 如果没有足够的 5 元找零，则返回 false
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			// 如果是 20 元，尽量找零 10 元和 5 元
			// 如果没有足够的 10 元找零，则找零 5 元
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	// 如果成功为每个顾客找零，则返回 true
	return true
}
