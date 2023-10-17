package main // 定义主程序包

/*import (
	"fmt"
	"math/rand"
	"time"
)

const ( // 定义常量
	n = 100000 // 数字总数
	d = 100    // 前d大的数字
)

// 快速排序算法
func quickSort(nums []int) { // 定义快速排序函数
	if len(nums) <= 1 {            // 如果只有一个数字，直接返回
		return
	}
	pivot := nums[len(nums)/2] // 获取中间的“轴点”
	i, j := 0, len(nums)-1     // 定义左右指针
	for i <= j {               // 如果左指针小于等于右指针
		for nums[i] > pivot { // 找到左边第一个小于等于pivot的数字
			i++
		}
		for nums[j] < pivot { // 找到右边第一个大于等于pivot的数字
			j--
		}
		if i <= j { // 如果左指针仍然小于等于右指针，交换左右指针指向的数字
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if j+1 >= d { // 如果右指针的位置比d小
		quickSort(nums[:j+1]) // 递归地对左半部分进行排序
	}
	if i-1 < n-d { // 如果左指针的位置比n-d大
		quickSort(nums[i:]) // 递归地对右半部分进行排序
	}
}

var (
	nums       [n + 1]int             // 存储数字的数组
	tree       [4*n + 1]int           // 存储二叉树的数组
	cmpCnt     int                    // 比较次数计数器
	winnerFunc = func(a, b int) int { // 使用函数变量封装比较操作
		cmpCnt++               // 计数器自增
		if nums[a] > nums[b] { // 比较a和b所表示的数字
			return a
		} else {
			return b
		}
	}
)

// 建立二叉树
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
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 随机化种子
	// 生成随机数，并进行去重
	m := make(map[int]bool) // 利用map来去除重复数字
	for i := 1; i <= n; {   // 逐个生成数字
		val := rand.Intn(1000000) // 生成一个从0到1000000的随机数
		if !m[val] {              // 如果这个数字没有出现过
			nums[i] = val // 将这个数字存入数组中
			m[val] = true // 表示这个数字已经出现过了
			i++           // 数组计数器自增
		}
	}
	//fmt.Println("Original array:", nums[1:]) // 打印原始数组
	// 排序
	quickSort(nums[0:])                    // 对nums[1:]中的n个数字进行排序
	fmt.Println("Sorted array:", nums[0:]) // 打印排序后的数字
	// 构建二叉树并排序
	buildTree(1, 1, n) // 建立二叉树
	divide(1, 1, n)    // 对二叉树进行分治排序

	// 输出比较次数
	fmt.Println("\nComparison Count:", cmpCnt)
}
*/
