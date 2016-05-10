##   AUPE (advanced programming in the unix environment)

### 编译步骤
* 下载源码
* 进入apue.3e 目录执行make 静态库文件在 apue.3e/lib/ 目录下
* 编译 编译时需要连接静态库文件 使用 _L 参数
```
 gcc -o test ls.c -L /Users/yaoyun/Downloads/apue3/apue.3e/lib -lapu
 _L 后面的目录使用libapue.a 所在目录
```

