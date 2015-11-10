package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []int
	var value reflect.Value = reflect.ValueOf(&a)
	// 判断指针是否指向内存地址
	if !value.CanSet() {
		value = value.Elem() //使指针指向内存地址
	}

	value = reflect.Append(value, reflect.ValueOf(1))
	value = reflect.Append(value, reflect.ValueOf(2))
	value = reflect.Append(value, reflect.ValueOf(3))
	value = reflect.Append(value, reflect.ValueOf(4), reflect.ValueOf(5)) //支持可变参数

	fmt.Println(value.Kind(), value.Slice(0, value.Len()).Interface())
	fmt.Printf("%+v", value)

}
