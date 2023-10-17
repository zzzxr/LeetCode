package main

/*
给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，
返回 0 。
输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
输出：16
解释：这两个单词为 "abcw", "xtfn"。
*/
func maxProduct(words []string) int {
	// 转ASCII码，再做偏移量
	asciiList := make([]int, len(words))
	for i, word := range words {
		for _, i3 := range word {
			asciiList[i] |= 1 << (i3 - 'a')
		}
	}
	// 比较是否存在重复字母并计算最大乘积
	area := 0
	for i, x := range asciiList {
		for j, y := range asciiList[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > area {
				area = len(words[i]) * len(words[j])
			}
		}
	}
	return area
}
