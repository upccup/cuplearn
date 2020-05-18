
* 多变量赋值时, 先计算所有相关值, 然后再从左到右依次赋值 
```
	data, i := [3]int{0,1,2,3}
	i, data[i] = 2, 100  //(i = 0) -> (i = 2),(data[0] = 100)
	fmt.Println(data)   //[100, 1, 2, 3]
	data[1], data[2] = data[2], data[1] // [100, 2, 1,3]
``` 

#### switch case 中的fallthrough 不同于 *break* 和 *continue* 不是在任何地方都可以使用的 
* The 'fallthrough' must be the last thing in the case; you can't write something like
```
switch {
case f():
    if g() {
        fallthrough // Does not work!
    }
    h()
default:
    error()
}
```
* However, you can work around this by using a 'labeled' fallthrough:
```
switch {
case f():
    if g() {
        goto nextCase // Works now!
    }
    h()
    break
nextCase:
    fallthrough
default:
    error()
}
```
#### swtich case 中的 fallthrough 会强制执行后面的代码
```
  integer := 6
  switch integer {
  case 4:
      fmt.Println("The integer was <= 4")
      fallthrough
  case 5:
      fmt.Println("The integer was <= 5")
      fallthrough
  case 6:
      fmt.Println("The integer was <= 6")
      fallthrough
  case 7:
      fmt.Println("The integer was <= 7")
      fallthrough
  case 8:
      fmt.Println("The integer was <= 8")
      fallthrough
  default:
      fmt.Println("default case")
  }
  这段程序的输出是
  The integer was <= 6
  The integer was <= 7
  The integer was <= 8
  default case
```

#### if 语句条件判断语句里面允许声明一个变量,但是这个变量的作用域只在该条件逻辑快内, 在其他地方就不起作用
```
  // 计算获取值x,然后根据x返回的大小，判断是否大于10。
  if x := computedValue(); x > 10 {
      fmt.Println("x is greater than 10")
  } else {
      fmt.Println("x is less than 10")
  }

  //这个地方如果这样调用就编译出错了，因为x是条件里面的变量
  fmt.Println(x)
```

### for select case 中如果使用break 只会break掉select而不会break掉外面的for

#### 秒的单位转换
Unit               | Symbol | Description
:-----------------:|:------:|:-------------------------------------:
nanosecond         |   ns   | 1 second = 1,000,000,000 nanoseconds
microsecond        |   μs   | 1 second = 1,000,000 microseconds
millisecond        |   ms   | 1 second = 1,000 milliseconds
second             |sec or s| base unit of Time
minute             |   min  | 1 minute = 60 seconds
hour               |   hr   | 1 hours = 60 minutes
day                |   d    | 1 day = 24 hours
week               |   wk   | 1 week = 7 days
fortnight          | 4tnite | 1 4tnite = 2 weeks or 14 days

#### os file 参数说明
 参数          |        值        | 说明
 :------------:|:----------------:|:--------------------------------------:
 O_RDONLY      |syscall.O_RDONLY  | open the file read-only.
 O_WRONLY      |syscall.O_WRONLY  | open the file write-only.
 O_RDWR        |syscall.O_RDWR    | open the file read-write.
 O_APPEND      |syscall.O_APPEND  | append data to the file when writing.
 O_CREATE      |syscall.O_CREAT   | create a new file if none exists.
 O_EXCL        |syscall.O_EXCL    | used with O_CREATE, file must not exist
 O_SYNC        |syscall.O_SYNC    | open for synchronous I/O.
 O_TRUNC       |syscall.O_TRUNC   | if possible, truncate file when opened.


### 关于数组和slice
  * array 就是数组,通过 var arr [n]type 来定义, 在 [n]type 中,n表示数组的长度,
type表示存储元素的类型.对数组的操作也是通过[]来进行读取和赋值.
  由于长度也是数组类型的一部分, 因此 [3]int 与 [4]int 是不同的类型,数组也就不能改变长度. 数组之间的赋值是值的赋值,
即当把一个数组作为参数传入函数的时候, 传入的其实是该数组的副本,而不是它的指针. 如果需要传入指针则需要用到slice.
  * slice 并不是真正意义上的动态数组, 而是一个引用类型. slice 总是指向一个底层 array, slice的声明也可以像 array  一样,只是不需要长度.
  slice 是引用类型, 所以当引用改变其中元素的值时,其他引用都会改变改值. 从概念上来说 slice 像一个结构体, 这个结构体包含了三个元素:一个指针 -- 指向数组中 slice 指定的开始位置, 长度-- 即 slice 的长度
最大长度 -- 也就是 slice 开始位置到数据的最后位置的长度.
  * channel，slice，map 都是引用类型, 作为参数传递是传值就可以改变内容, 但是如果需要改变 slice 的长度仍然需要传地址.


### 方法接收者
如果方法的接收者为对象的指针,则可以修改原对象.如果方法接收者为对象的值,那么方法中备操作的是原对象的副本, 对其操作不会影响原对象.

go语言中函数的参数默认为 **按值传递**, 即在函数内部修改传入参数的值是外部传入值的拷贝, 对象的方法也是遵循这个逻辑, 如果方法的接收者为对象指针, 其实函数传递的也是指针的拷贝,但是指针的拷贝也还是指向原有对象.

如果对象本身是引用类型, 例如: slice, map.则即使方法的接收是对象的值也是可以对其进行修改的.

* 注: java语言中函数也是传值的但是java中大部分对象都是引用类型的,所以可以在函数中对其值进行修改.
* 当我们使用指针实现接口时，只有指针类型的变量才会实现该接口；当我们使用结构体实现接口时，指针类型和结构体类型都会实现该接口, 所以判断某个对象是否实现了某个接口时,如果该方法是指针类型实现,则结构体类型在类型推断时不能通过.

[测试代码](https://github.com/upccup/cuplearn/blob/master/go-study/builtin/function/func.go)


### 接口
Go 语言中接口是**隐式实现**的,只要实现了接口中所定义的所有方法,则表示实现了该接口.
接口也是Go 语言中的一种类, 它能够出现在变量的定义,函数的入参和返回值中并对他们进行约束.
Go 语言中有两种略微不同的接口,一种是带有一组方法的接口,另一种是不带任何方法的interface{}. 需要注意的是与C语言中的void * 不同, interface{}类型不是**任意类型**,变量在转换成interface{}类型后仍然带有原类型的信息. [示例代码](https://github.com/upccup/cuplearn/blob/master/go-study/interface/eface.go)
```
    package main

    type TestStruct struct{}

    func NilOrNot(v interface{}) bool {
        return v == nil
    }

    func main() {
        var s *TestStruct
        fmt.Println(s == nil)      // #=> true
        fmt.Println(NilOrNot(s))   // #=> false
    }

    $ go run main.go
    true
    false
```
出现上述现象的原因是 —— 调用 NilOrNot 函数时发生了隐式的类型转换，除了向方法传入参数之外，变量的赋值也会触发隐式类型转换。在类型转换时，*TestStruct 类型会转换成 interface{} 类型，转换后的变量不仅包含转换前的变量，还包含变量的类型信息 TestStruct，所以转换后的变量与 nil 不相等.

[参考文档](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/)

### 位运算

* & 与 AND 只有两个数都是1结果才为1
* | 或 OR 两个数有一个是1 结果就是1
* ^ 异或 XOR ^作二元运算符就是异或，相同为0，不相同为1,  ^作一元运算符表示是按位取反(一个有符号位的^操作为 这个数+1的相反数)
* &^ 位与非 AND NOT。它是一个 使用 AND 后，再使用 NOT 操作的简写 将运算符左边数据相异的位保留，相同位清零
* << 左移 左移规则:  右边空出的位用0填补, 高位左移溢出则舍弃该高位
* >> 右移 右移规则：左边空出的位用0或者1填补。正数用0填补，负数用1填补。注：不同的环境填补方式可能不同, 低位右移溢出则舍弃该位

[代码示例](https://github.com/upccup/cuplearn/blob/master/go-study/builtin/bit-operation.go)

操作符运算的优先级
| 优先级 |    符号 
:------:|:---------:
1       | ` *   /   %   <<  >>  &   &^ `
2       | ` +   -   |   ^ `
3       | ` ==  !=  <   <=  >   >= `
4       | ` && `
5       | ` || `

* 反码: 正数的反码等于它的原码, 负数的反码则保留其原码符号位，然后对其他位进行取反操作
* 补码: 正数的补码是其本身, 负数的补码是在它的反码基础上加 1

补码和反码的详细介绍参考[深入浅出: 解读原码、反码和补码](https://juejin.im/post/5ddce6dee51d4532ca58e5e6)
