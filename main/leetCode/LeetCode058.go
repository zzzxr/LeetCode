package main

/*
最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为4。
*/

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	length := 0
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		length++
		index--
	}
	return length
}
