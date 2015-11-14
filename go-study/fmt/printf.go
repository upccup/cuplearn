package main

import (
	"fmt"
)

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

type T struct {
	a int
	b float32
	c string
}

func main() {
	t := &T{7, -2.35, "test\tpk"}
	// %v      相应值的默认格式
	fmt.Printf("NO1 %s: %v\n", "%v", t)
	// NO1 %v: &{7 -2.35 test	pk}

	// %+v     打印结构体时，会添加字段名
	fmt.Printf("NO2 %s: %+v\n", "%+v", t)
	// NO2 %+v: &{a:7 b:-2.35 c:test	pk}

	// %#v     相应值的Go语法表示
	fmt.Printf("NO3 %s: %#v\n", "%#v", t)
	// NO3 %#v: &main.T{a:7, b:-2.35, c:"test\tpk"}
	fmt.Printf("NO4 %s: %#v\n", "%#v", timeZone)
	// NO4 %#v: map[string]int{"CST":-21600, "MST":-25200, "PST":-28800, "UTC":0, "EST":-18000}

	// %T      相应值的类型的Go语法表示
	fmt.Printf("NO5 %s: %T\n", "%T", t)
	// NO5 %T: *main.T

	var yes bool //初始化默认值为false
	// %t  值的true 和 false
	fmt.Printf("NO6 %s: %t\n", "%t", yes)
	// NO6 %t: false

	// %T      相应值的类型的Go语法表示
	fmt.Printf("NO7 %s: %T\n", "%t", yes)
	// NO7 %t: bool

	// %% 打印一个 % 号
	fmt.Printf("NO8 %s: %%\n", "%%")
	// NO8 %%: %

	// %q 给打印的字串自动加引号
	fmt.Printf("NO9 %s: %q\n", "%q", "yaoyuntest")
	// NO9 %q: "yaoyuntest"

	// **************** 整数 ********************

	// 二进制表示
	fmt.Printf("No10: %s %b\n", "%b", 888)
	// No10: %b 1101111000

	//%c 数值对应的Unicode编码字符
	fmt.Printf("NO11: %s %c\n", "%c", '姚') //单引号
	fmt.Printf("NO11: %s %c\n", "%c", "姚") //双引号
	// NO11: %c 姚
	// NO11: %c %!c(string=姚)

	// %d 十进制表示
	fmt.Printf("NO12: %s %d\n", "%d", 888)
	// NO12: %d 888

	// %o 八进制表示
	fmt.Printf("NO13: %s %o\n", "%o", 888)
	// NO13: %o 1570

	// 十六进制表示，使用a-f
	fmt.Printf("NO14: %s %x\n", "%x", 8881)
	// NO14: %x 22b1

	// 十六进制表示，使用A-F
	fmt.Printf("NO15: %s %X\n", "%X", 8881)
	// NO15: %X 22B1

	// %U Unicode格式： U+1234，等价于"U+%04X"
	fmt.Printf("NO16: %s %U\n", "%U", 888)
	// NO16: %U U+0378

	// %q 单引号
	fmt.Printf("NO17 %s %q\n", "%q", 888)
	// NO17 %q '\u0378'

	//********  浮点数 *******************

	// %b 无小数部分、两位指数的科学计数法，和strconv.FormatFloat的'b'转换格式一致。
	fmt.Printf("NO18 %s %b\n", "%b", 888.66)
	// NO18 %b 7816736025115361p-43

	// 对于 %g/%G 而言，精度为所有数字的总数，例如：123.45，%.4g 会打印123.5，（而 %6.2f 会打印123.45）。
	// %e 和 %f 的默认精度为6
	// %e 科学计数法，举例：-1234.456e+78
	fmt.Printf("NO19: %s %e\n", "%e", 888.66)
	// NO19: %e 8.886600e+02

	// %E 科学计数法，举例：-1234.456E+78
	fmt.Printf("NO19: %s %E\n", "%E", 888.66)
	// NO19: %E 8.886600E+02

	// %f 有小数部分，但无指数部分，举例：123.456
	fmt.Printf("N020: %s %f\n", "%f", 888.66)
	// N020: %f 888.660000

	// %g 根据实际情况采用%e或%f格式（以获得更简洁的输出）
	fmt.Printf("NO21: %s %g\n", "%g", 888.66)
	// NO21: %g 888.66
	fmt.Printf("NO21: %s %g\n", "%g", 99999999999.93)
	// NO21: %g 9.999999999993e+10

	// %G 根据实际情况采用%E或%f格式（以获得更简洁的输出）
	fmt.Printf("NO22: %s %G\n", "%G", 888.66)
	// NO22: %G 888.66
	fmt.Printf("NO22: %s %G\n", "%G", 99999999999.93)
	// NO22: %G 9.999999999993E+10

	// "%.*d",10,123: 打印整数，并保证10为长度，关键是长度运行时动态传入
	fmt.Printf("NO23: %s %.*d\n", "%.*d", 10, 123)
	// NO23: %.*d 0000000123

	// %6.*f",2,888.666: 打印浮点数，并保证两位小数，关键是小数位运行时动态传入
	fmt.Printf("NO24: %s %6.*f\n", "%6.*f", 2, 888.666)
	// NO24: %6.*f 888.67

}
