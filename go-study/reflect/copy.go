package main

import (
	"fmt"
	"reflect"
)

// func Copy(dst, src Value) int
// 参数列表: dst Value 是目标切片Slice 或数组Array  src Value 是源切片Slice 或数组Array
// 返回值: 返回 int 复制过去元素的数量
// 功能说明: Copy 复制src的内容复制到dst, 直到dst一杯填满或src已被耗尽. 它返回复制的元素的数量. 每个dist和src 的Kind(样)
// 都必须切片(Slice)或数组(Array), dst和src必须具有相同的元素类型

type A struct {
	A0 []int
	A1 []int
}

func main() {
	var a A
	a.A0 = append(a.A0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	a.A1 = append(a.A1, 9, 8, 7, 6)
	var n = reflect.Copy(reflect.ValueOf(a.A0), reflect.ValueOf(a.A1))
	fmt.Println(n, a) // 4 {[9 8 7 6 5 6 7 8 9] [9 8 7 6]}
}
