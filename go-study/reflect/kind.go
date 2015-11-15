package main

import (
	"fmt"
	"reflect"
)

// 功能说明

/* ***************************
一个类型代表一种特定类型的类型，零Kind不是有效的Kind
reflect.Kind 有以下常量成员
reflect.Invalid       // 无效
reflect.Bool          // 布尔
reflect.Int           // 整数（有符号）
reflect.Int8          // 整数8位（有符号）
reflect.Int16         // 整数16位（有符号）
reflect.Int32         // 整数32位（有符号）
reflect.Int64         // 整数64位（有符号）
reflect.Uint          // 整数（无符号）
reflect.Uint8         // 整数8（无符号）
reflect.Uint16        // 整数16（无符号）
reflect.Uint32        // 整数（无符号）
reflect.Uint64        // 整数（无符号）
reflect.Uintptr       // 整数（指针,无符号）
reflect.Float32       // 浮点数32位
reflect.Float64       // 浮点数64位
reflect.Complex64     // 复数64位
reflect.Complex128    // 复数128位
reflect.Array         // 数组
reflect.Chan          // 信道
reflect.Func          // 函数
reflect.Interface     // 接口
reflect.Map           // 地图
reflect.Ptr           // 指针
reflect.Slice         // 切片
reflect.String        // 字符
reflect.Struct        // 结构
reflect.UnsafePointer // 安全指针
******************************/

func main() {
	var a string
	var kind reflect.Kind = reflect.TypeOf(a).Kind()
	var kind1 reflect.Kind = reflect.ValueOf(a).Kind()
	fmt.Println(kind, kind == reflect.String, kind == reflect.Int)    // string true false
	fmt.Println(kind1, kind1 == reflect.String, kind1 == reflect.Int) // string true false
}
