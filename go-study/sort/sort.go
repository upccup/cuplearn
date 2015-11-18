package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// func Sort(data Interface)
//参数列表: data 表示要排序的 Interface 数据
// 功能说明 Sort 对date进行排序. 它调用一次 dara.Len 来决定排序的长度
// 调用 dara.Less和data.Swap 的开销为 O(n*log(n)) 此排序为不稳定排序

// 代码实现
/*
type Interface interface {
	// Len is the nuber of elements in the collection
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j
	Less(i, j int) bool
	// Swap swaps the elements woth indexex i and j
	Swap(i, j int)
}
*/

// 任何实现了 sort.Interface 的类型(一般为集合), 均可使用该包的方法进行排序
// 这这些方法要求集合内列出元素的索引为整数

type ByTime []os.FileInfo

func (t ByTime) Len() int {
	return len(t)
}

func (t ByTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByTime) Less(i, j int) bool {
	return t[i].ModTime().Before(t[j].ModTime())
}

type ByLenth []string

func (s ByLenth) Len() int {
	return len(s)
}

func (s ByLenth) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLenth) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	files, _ := ioutil.ReadDir("../")
	sort.Sort(ByTime(files))
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i].ModTime())
		// 2015-08-10 14:43:35 +0800 CST
		// 2015-08-10 16:03:25 +0800 CST
		// ...
	}

	d := []int{5, 2, 4, 6, 5, 5, 1, 11, 43, 0, -1}
	sort.Sort(sort.IntSlice(d))
	fmt.Println(d) // [-1 0 1 2 4 5 5 5 6 11 43]

	a := []float64{1.2, -1, 0.0, 1.0, 89.2, 0.0001, 1.2345}
	sort.Sort(sort.Float64Slice(a))
	fmt.Println(a) // [-1 0 0.0001 1 1.2 1.2345 89.2]

	s := []string{"PHP", "go", "Go", "java", "python", "c"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s) // [Go PHP c go java python]

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLenth(fruits))
	fmt.Println(fruits) // [banana kiwi peach]
}
