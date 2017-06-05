|   连接词            |           意义       |
|:-------------------:|:--------------------:|
| 文件表达式
| if [ -f file ]      |  如果文件存在
| if [ -d ... ]       |  如果目录存在
| if [ -s file ]      |  如果文件存在且非空
| if [ -f file ]      |  如果文件存在且可读
| if [ -w file ]      |  如果文件存在且可写
| if [ -x file ]      |  如果文件存在且可执行
| 整数变量表达式
| if [ int1 -eq int2 ] |  如果两整数相等
| if [ int1 -ne int2 ] |  如果两整数不相等
| if [ int1 -ge int2 ] |  如果 大于等于
| if [ int1 -gt int2 ] |  如果 大于
| if [ int1 -le int2 ] |  如果 小于等于
| if [ int1 -lt int2 ] |  如果 小于
| 字符串变量表达式
| if [ $string1 = $string2 ]       |  如果 字符串相等
| if [ $string1 != $ string2 ]     |  如果 字符串不相等
| if [ -n $string ]                |  如果 字符串 非空(非0)
| if [ -z $string ]                |  如果 如果字符串为空
| if [ &string ]                   |  如果  字符串非空 (同 -n)
| 逻辑表达式
| if [ ! 表达式 ]                  |  逻辑非
| if [ 表达式1 -a 表达式2 ]        |  逻辑与
| if [ 表达式1 -o 表达式2 ]        |  逻辑或
| if简化语句
| && 如果是“前面”，则“后面”        | [ -f /var/run/dhcpd.pid ] && rm /var/run/dhcpd.pid    检查 文件是否存在，如果存在就删掉
|  \|\|   如果不是“前面”，则后面     | [ -f /usr/sbin/dhcpd ] \|\| exit 0    检验文件是否存在，如果存在就退出


* 注: **=** 用来判断等于时 等号两边必须都有空格, 等号用来赋值时正好相反,在等号两边不能有空格