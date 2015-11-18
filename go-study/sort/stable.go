package main

import (
	"fmt"
	"sort"
)

// func Stable(data Interface)
// 参数列表: data表示要排序的Interface数据
// 功能说明: Stable 稳定排序算法, 算法将会将相等的元素维持其相对次序. 如果一个排序算法也是稳定的
//          当有两个相等的元素之R和S, 而且原本的列表中R出现在S之前, 那么排序过的列表中R也将会是在S之前.
//          对于比较排序算法,我们都能给出n个输入的数值, 使算法以Ω(n*logn)时间运行.
// 			稳定排序算法: 插入排序, 冒泡排序, 归并排序, 计数排序, 基数排序, 桶排序

type MyString []string

func (s MyString) Len() int {
	return len(s)
}

func (s MyString) Less(i, j int) bool {
	if s[i] == "" {
		return true
	}

	if s[j] == "" {
		return false
	}

	return []byte(s[i])[0] < []byte(s[j])[0] //比较字符串的第一个字符大小
}

func (s MyString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	data := MyString{"5A*", "24", "65", "23", "1", "57", "4", "624"} // unsorted
	sort.Stable(data)
	fmt.Println(data) // [1 24 23 4 5A* 57 65 624]
}
