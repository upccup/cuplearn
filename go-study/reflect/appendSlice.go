package main

import (
	"fmt"
	"reflect"
)

// 参数列表
// s Value 原切片数据, 类型是reflect.Value
// t Value 新切片数据, 将此切片数据追加到s切片中 类型是reflect.Value

// 返回值
// Value 返回新的切片

// 功能说明
// 追加一个切片t到另一个切片s,并返回所创建的Slice, 在Go中, 没一个x值必须是分配给切片的元素类型

func main() {
	var a []int
	var value reflect.Value = reflect.ValueOf(&a)

	//判断指针是否指向内存地址
	if !value.CanSet() {
		value = value.Elem() //使指针指向内存地址
	}

	value = reflect.AppendSlice(value, reflect.ValueOf([]int{1, 2}))                //支持切片
	value = reflect.AppendSlice(value, reflect.ValueOf([]int{3, 4, 5, 6, 7, 8, 9})) //支持切片
	fmt.Println(value.Kind(), value.Slice(0, value.Len()).Interface())
	/////  >> slice [1 2 3 4 5 6 7 8 9]

}
