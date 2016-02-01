
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
