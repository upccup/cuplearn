## tr

#### tr命令可以对来自标准输入的字符进行替换, 压缩和删除. 可以把 tr 当做 sed 的简化版本.

#### 语法
```
tr [OPTION]…SET1[SET2]  
```

#### 选项说明
* -c, --complement: 反选设定字符, 也就是符合SET1的部分不做处理, 不符合的剩余部分才进行转换
* -d, --delete: 删除指定字符
* -s, --squeeze-repeats: 缩减连续重复的字符成指定的单个字符
* -t, --truncate-set1: 削减SET1使其与SET2的长度相同
* --help：显示程序用法信息
* --version：显示程序本身的版本信息

#### 参数说明
* 字符集1：指定要转换或删除的原字符集。当执行转换操作时，必须使用参数“字符集2”指定转换的目标字符集。但执行删除操作时，不需要参数“字符集2”；
* 字符集2：指定要转换成的目标字符集

#### 实例
* 将输入字符大写转为小写
```
echo "HELLO WORLD" | tr 'A-Z' 'a-z'
hello world
```
注: 'A-Z' 和 'a-z'都是集合，集合是可以自己制定的，例如：'ABD-}'、'bB.,'、'a-de-h'、'a-c0-9'都属于集合，集合里可以使用'\n'、'\t'，可以可以使用其他ASCII字符。

* 使用tr删除字符：
```
echo "hello 123 world 456" | tr -d '0-9' 
hello world 
```

* 将制表符转换为空格
```
cat text | tr '\t' ' '
```

* 删除指定的字符
```
echo 'aa.,a 1 b#$bb 2 c*/cc 3 ddd 4' | tr -d -c '0-9 \n'
 1  2  3  4
```

* 压缩字符, 删除重复的字符
```
echo "thissss is a text linnnnnnne." | tr -s ' sn'
this is a text line.
```

#### tr 可以使用的字符类
```
\NNN 八进制值的字符 NNN (1 to 3 为八进制值的字符)
\\ 反斜杠
\a Ctrl-G 铃声
\b Ctrl-H 退格符
\f Ctrl-L 走行换页
\n Ctrl-J 新行
\r Ctrl-M 回车
\t Ctrl-I tab键
\v Ctrl-X 水平制表符
CHAR1-CHAR2 ：字符范围从 CHAR1 到 CHAR2 的指定，范围的指定以 ASCII 码的次序为基础，只能由小到大，不能由大到小。
[CHAR*] ：这是 SET2 专用的设定，功能是重复指定的字符到与 SET1 相同长度为止
[CHAR*REPEAT] ：这也是 SET2 专用的设定，功能是重复指定的字符到设定的 REPEAT 次数为止(REPEAT 的数字采 8 进位制计算，以 0 为开始)
[:alnum:] ：所有字母字符与数字
[:alpha:] ：所有字母字符
[:blank:] ：所有水平空格
[:cntrl:] ：所有控制字符
[:digit:] ：所有数字
[:graph:] ：所有可打印的字符(不包含空格符)
[:lower:] ：所有小写字母
[:print:] ：所有可打印的字符(包含空格符)
[:punct:] ：所有标点字符
[:space:] ：所有水平与垂直空格符
[:upper:] ：所有大写字母
[:xdigit:] ：所有 16 进位制的数字
[=CHAR=] ：所有符合指定的字符(等号里的 CHAR，代表你可自订的字符)
```

使用方法: 
```
echo "abc Asdd11" | tr '[:lower:]' '[:upper:]'
ABC ASDD11
```