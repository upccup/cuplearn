
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
