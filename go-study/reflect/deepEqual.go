package main

import (
	"fmt"
	"reflect"
)

// func DeepEqual(a1, a2 interface{}) bool
// 参数列表: 待比较的两个对象a1, a2
// 返回值: bool 如果两个参数的类型和值都是相等的返回true 否则返回false
// 功能说明: DeepEqual 在可能的情况下默认使用 ==  来判断相等, 将扫描Array, Slice, Map, Struct和字段	(Filed)的成员
// 正确处理递归类型. 如果两个参数都是nil 则返回true

func main() {
	equalInt()
	equalStruct()
}

func equalInt() {
	var a, b int
	a = 1
	b = 2
	fmt.Println(reflect.DeepEqual(a, b)) // false
	b = 1
	fmt.Println(reflect.DeepEqual(a, b)) // true
}

func equalStruct() {
	type A struct {
		A0 int
	}

	type B struct {
		A0 int
	}

	var (
		a A
		b B
	)

	fmt.Println(reflect.DeepEqual(a, b))       // false
	fmt.Println(reflect.DeepEqual(a.A0, b.A0)) // true
}
