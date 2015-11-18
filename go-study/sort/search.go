package main

import (
	"fmt"
	"sort"
)

// func Search(n int, f func(int) bool) int
// 参数列表: n 表示切片的长度 f func(int) bool 类型的函数
// 返回值: 返回切片的索引
// 功能说明: Search 使用二分查找法在[0,n) 中寻找并返回 f(i)==true 的最小索引, 假设该索引
// 在区间[0,N) 内则f(i)==true蕴含f(i+1)==true. Search 要求f对于输入区间[0,n)(可能为空)的前一部分为false
// 而对于剩余(可能为空)的部分为true
// Search 返回第一个f为true时的索引i. 若该索引不存在, Search 就返回n. Search 仅当i早区间[0,n)内时才调用f(i)
// Search 常用于在一个已排序的,可索引的数据结构中寻找索引为i的值z, 例如数组或切片.这种情况下,实参f,一般是一个闭包
// 会捕获所要搜索的值,以及索引并排序该数据结构的方式
// 例如, 给定一个以升序排列的切片数据, 调用
/*
	Search(len(data), func(i int) bool { retrurn data[i] >= 23 })
*/
// 会返回满足data[i]>=23的最小索引i.若调用者想要判断23是否在切片中,就必须单独测试data[i] == 23的值
// 搜索降以序排列的数据, 需使用 <= 操作符, 而非 >= 操作符

type finder struct {
	data []int
	targ int
	f    func(n int) bool
}

func MakeFinder1() *finder {
	var f finder
	f.f = func(i int) bool {
		return f.data[i] >= f.targ
	}

	return &f
}

func MakeFinder2() *finder {
	var f finder
	f.f = func(i int) bool {
		return f.data[i] <= f.targ
	}

	return &f
}

func (f *finder) Find(data []int, x int) int {
	f.data = data
	f.targ = x
	return sort.Search(len(f.data), f.f)
}

func main() {
	var data1 = []int{9, 19, 29, 39, 49, 59, 69, 79, 89, 99}
	var data2 = []int{99, 89, 79, 69, 59, 49, 39, 29, 19, 9}
	f1 := MakeFinder1()
	i := f1.Find(data1, 50)
	fmt.Println(i, data1[i]) // 5 59

	f2 := MakeFinder2()
	i = f2.Find(data2, 40)
	fmt.Println(i, data2[i]) // 6 39
	if i < len(data2) && f2.targ == data2[i] {
		fmt.Printf("fidn %v in data2\n", f2.targ)
	} else {
		fmt.Printf("can't find %v in data2\n", f2.targ)
	} //can't find 40 in data2
}

// 源码
func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
