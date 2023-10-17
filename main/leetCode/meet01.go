package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
多路快排
在10万个数字中寻找前 100大 (top 100)的数字，要求在20万次比较中得到答案
https://blog.csdn.net/weixin_44137652/article/details/130977706 要求相同， 但是该解法错误
*/
const (
	n = 100
	d = 10
)

var (
	nums       [n + 1]int             // 存储数字的数组
	tree       [4*n + 1]int           // 存储二叉树的数组
	count      int                    // 比较次数计数器
	winnerFunc = func(a, b int) int { // 使用函数变量封装比较操作
		count++                // 计数器自增
		if nums[a] > nums[b] { // 比较a和b所表示的数字
			return a
		} else {
			return b
		}
	}
)

func buildTree(node, l, r int) { // 定义建立二叉树的函数
	if l == r { // 如果只剩下一个数字
		tree[node] = l // 将该数字作为当前节点
		return
	}
	mid := (l + r) / 2                                    // 获取区间的中间位置
	buildTree(node*2, l, mid)                             // 递归地建立左子树
	buildTree(node*2+1, mid+1, r)                         // 递归地建立右子树
	tree[node] = winnerFunc(tree[node*2], tree[node*2+1]) // 将这两个子节点中的胜者作为当前节点
}

// 分治算法，对二叉树的节点进行排序
func divide(node, l, r int) { // 定义分治排序函数，对二叉树的节点进行排序
	// 叶子节点
	if l == r { // 如果只剩下一个数字
		if r <= d { // 如果该数字是前d大的数字之一，则将它输出
			fmt.Print(nums[l], " ")
			nums[l] = 0
		}
		return
	}
	// 非叶子节点
	mid := (l + r) / 2                                    // 获取当前区间的中间位置
	divide(node*2, l, mid)                                // 递归处理左子树，得到左子树中前100大的数字
	divide(node*2+1, mid+1, r)                            // 递归处理右子树，得到右子树中前100大的数字
	tree[node] = winnerFunc(tree[node*2], tree[node*2+1]) // 将这两个子节点中的胜者作为当前节点
	// 输出前d大数字
	if l <= 1 && r >= d {
		fmt.Print(nums[tree[node]], " ")
		nums[tree[node]] = 0
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	numMap := make(map[int]bool)
	for i := 0; i < n; {
		num := rand.Intn(n)
		if !numMap[num] {
			nums[i] = num
			numMap[num] = true
			i++
		}
	}
	fmt.Println("原始数组：", nums)
	buildTree(0, 0, len(nums)-1)
	divide(0, 0, len(nums)-1)
	fmt.Println("排序次数：", count)
}
