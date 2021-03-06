### sprintf 格式化字符串函数说明
#### sprintf(格式化之后的字符串, "%(长度)格式符", 格式化之前的字符串)
1. d 格式符. 用来输出10进制数
   * %d 安装整数的实际长度输出.
   * %md m为指定的输出字段的宽度. 如果数据小于m,则左端以空格补齐, 若大于m, 则按实际长度输出.
   *  %ld 输出长整型数据. 例如: long a = 125790; printf(%ld", a); 如果用%d输出就会发生错误, 因为整型数据范围是 -32768到32768. 对long型数据应该用%ld格式输出, 对长整型数据也可以指定字符宽度, 如: %8ld.
2. o格式符, 以八进制数形式输出整数. 由于内存单元中的各位的值(0或1)按八进制输出,因此输出的数值不带符号, 即将符号位也一起作为八进制的一部分输出
3. x格式符 以十六进制形式输出整数.同样不会出现负的十六进制数.同样可以用%lx输出长整型数,也可以指定输出字段宽度.
4. u格式符 用来输出 unsigned 型数据, 即无符号数, 以十进制形式输出.一个有符号整数(int型)也可以用 %u 格式输出, 反之, 一个 unsigned 型的数据也可以用 %d 输出. 按相互赋值的规则处理. unsigned 型数据也可以用 %x 格式输出.
5. c格式符, 用来输出一个字符, 一个字符型数据也可以用整数形式输出.
6. s格式符, 用来输出一个字符串.
   * %s printf("%s", "HELLO")
   * %ms 输出字符粘m列, 如果字符本身大于m, 则突破m限制字符全部输出, 如果小于m, 则左补齐空格.
   * %-ms 如果字符串小于m, 则在m范围内, 字符串向左靠, 右补齐空格
   * %m.ns 输出占m列, 但只取字符串左端n个字符. 这n个字符输出在m列的右侧, 左补齐空格
   * %-m.ns 其中m, n含义同上, n个字符输出在m范围的左侧, 右补齐空格. 如果 n>m 则自动取n值, 即保证n个字符正常输出.
7. f格式符用来输出实数(包括单,双精度), 以小数形式输出.
   * %f 不指定字符宽度, 由系统自动指定, 使整数部分全部如数输出,并输出6位小数. 需要注意并非全部数字都是有效数字,单精度实数的有效数字一般是7位
   * %m.nf 指定输出的数据占m列, 其中有n位小数.如果数值长度小于m, 则左补齐空格
   * %-m.nf 与 %m.nf 基本相同, 只是输出的数值向左端靠齐, 右端补空格.
8. e格式符, 以指数形式输出.且使用小写的e. E 格式符使用E.
9. g格式符, 用来输出实数, 它个悲剧数值的大小,自动选择f格式或e格式(选择输出时占宽度较小的一种), 且不输出无意义的零.G则输出E与f中较短的一个.

##### 说明
* 一个位于一个%和格式化命令间的整数担当者一个最小字段宽度说明符, 并且加上足够多的空格或0使输出自够长.如果想填充0,在最小宽度说明符前放置0.
