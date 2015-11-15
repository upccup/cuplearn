package main

import (
	"fmt"
	"reflect"
)

// func Indirect(v Value) Value
//参数列表: v Value 传入的是一个reflect.Value 类型的变量
//返回值: Value输出v的指针
//功能说明: 返回v的指针 如果v是一个零指针,简介返回零值 如果v不是一个指针, 间接的返回v

func main() {
	var a []int

	var value reflect.Value = reflect.ValueOf(&a)
	var value1 reflect.Value = reflect.ValueOf(a)

	fmt.Println(value.Kind())  // ptr
	fmt.Println(value1.Kind()) // slice

	// 判断指针是否指向内存
	if !value.CanSet() {
		fmt.Println("value pointer does not point to memory")
		value = value.Elem() //使指针指向内存
	}

	if !value1.CanSet() {
		fmt.Println("value1 pointer does not point to memory")
		// value1 = value1.Elem()   panic: reflect: call of reflect.Value.Elem on slice Value
	}
	fmt.Println(value.Kind()) // slice

	var value2 reflect.Value = reflect.ValueOf(a)
	var value3 reflect.Value = reflect.ValueOf(&a)
	value2 = reflect.Indirect(value3)
	fmt.Println(value2.Kind()) // slcie
	value3 = reflect.Indirect(value3)
	fmt.Println(value3.Kind()) // slice

}

// Indirect returns the value that v points to.
// If v is a nil pointer, Indirect returns a zero Value.
// If v is not a pointer, Indirect returns v.
/* **********************
func Indirect(v Value) Value {
	if v.Kind() != Ptr {
		return v
	}
	return v.Elem()
}
*************************/
