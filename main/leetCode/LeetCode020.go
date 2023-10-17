package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

var valueMap = map[uint8]uint8{')': '(', '}': '{', ']': '['}

func isValid(s string) bool {
	// 定义一个栈，讲字符串的值按照顺序放入栈中，在放入之前比较，如果成对，则不放入，并且将栈顶的值弹出，否则放入栈顶
	// 没有现成的栈，还是用数组实现
	if len(s)%2 != 0 {
		return false
	}
	temp := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if len(temp) == 0 {
			temp = append(temp, s[i])
			continue
		}
		if v, ok := valueMap[s[i]]; ok && v == temp[len(temp)-1] {
			temp = temp[:len(temp)-1]
		} else {
			temp = append(temp, s[i])
		}
	}
	return len(temp) == 0
}
