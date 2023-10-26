package main

/*
无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {
	// 定义长度
	length := len(s)
	// 判断如果长度小于2，直接返回长度即可
	if length < 2 {
		return length
	}
	// 定义左指针，右指针，最大子串长度
	left, right, maxLength := 0, 0, 0
	// 定义存储字符的map
	byteMap := map[byte]bool{}

	// 移动指针判断
	for right < length {
		if _, ok := byteMap[s[right]]; ok {
			// 如果右指针所在的字符已经存在于map中，则删除左指针的字符，并右移左指针
			delete(byteMap, s[left])
			left++
		} else {
			// 否则，保存右指针字符到map中，并计算子串长度，同时右移右指针
			byteMap[s[right]] = true
			maxLength = max(maxLength, right-left+1)
			right++
		}
	}
	return maxLength
}
