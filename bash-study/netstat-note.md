### netstat 使用说明

#### 连接状态码说明
* LISTEN: 侦听来自远方的TCP端口的连接请求
* SYN-SENT: 再发送连接请求后等待匹配的连接请求
* SYN-RECEIVED: 再收到和发送一个连接请求后等待对方连接请求的确认
* ESTABLISHED: 代表一个打开的连接
* FIN-WAIT-1: 等待远程TCP连接中断请求, 或先前的连接中断请求的确认
* FIN-WAIT-2: 从远程TCP等待连接中断请求
* CLOSE-WAIT: 等待本地用户发来的连接中断请求
* CLOSING: 等待远程TCP对连接中断的确认
* LAST-ACK: 等待原来的发向远程TCP的连接中断请求的确认
* TIME-WAIT: 等待足够的时间以确保远程TCP接收到连接中断请求的确认
* CLOSED: 没有任何连接状态

#### 参数说明
* 列出所有连接
```
  列出所有当前的连接使用  -a   选项即可
```
* 只列出 TCP 或 UDP 协议的连接
```
  使用  -t   选项列出TCP协议的连接
  使用  -u   选项列出UDP协议的连接
```
* 禁用反向域名机械, 加快插叙速度
```
  默认情况下 netstat 会通过反向域名解析技术查找每个IP对应的主机名,
这会降低查询速度, 如果觉得IP地址已经足够, 而没必要知道主机名,
就使用  -n   选项禁用域名解析
```
* 只列出监听中的链接
```
  任何网路服务的后台进程都会打开一个端口, 用于监听接入的请求,
这些正在监听的套接字也和连接的套接字一样, 也能被netstat列出来, 使用
-l   选项列出正在监听的套接字
  注意不要使用  -a  选项否则netstat会列出所有连接,
而不仅仅是监听的端口1
```
* 获取进程名, 进程号以及用户ID
```
  使用   -p   选项查看进程信息
  使用 -p 选项时 netstat 必须运行在root 权限之下,
不然它就得不到运行在root权限下的进程名
  使用   -ep   选项可以同时查看用户名和进程名
  如果将 -n  和  -e   一起使用 User
列的属性就是用户的ID号而不是用户名
```
* 打印统计数据
```
  使用  -s   可以打印出网络统计数据, 包括某个协议下的收发包数量
```
* 显示内核路由信息
```
  使用  -r   选项打印内核路由信息,
打印出来的信息与route命令输出的信息一样, 我们也可以使用  -n
选项禁止域名解析
```
* 打印网络接口
```
  使用  -i   选项能够打印网络接口信息
  把 -e 和 -i 一起使用 就相当于 ifconfig 了
```
* netstat 持续输出
```
  使用  -c  选项可以让 netstat 持续输出信息
```
* 显示多播组信息
```
  使用  -g   选项会输出 IPv4 和 IPv6 的多播组信息
```
* 打印 active 状态的连接
```
  active 状态的套接字连接用  ESTABLISHED 状态码表示, 所以可以使用 grep
命令活的active状态的连接
  配合 watch 监视active 状态的连接
  watch -d -n0 "netstat -atnp | grep ESTA"
```
* 查看服务是否正在运行
```
  如果想要查看服务是否正在运行使用 grep 例如查看ntp服务是否正在运行
  sudo netstat -aple | grep ntp
```



#### 常用命令说明
* 查看帮助信息
```
  netstat --help
```
* 显示当前所有活动的网络连接
```
  netstat -na
```
* 显示出所有处于监听状态的应用程序及进程号和端口号
```
  netstat -aulntp
```
* 显示所有80端口的网络连接
```
  netstat -aulntp | grep 80
```
* 如果还想对返回的连接列表进行排序. 就要用到sort命令了
```
  netstat -aulntp | grep 80 | sort
```
* 如果还想进行统计的话, 就可以往后面加wc命令
```
  netstat -aulntp | grep 80 | wc -l
```
* 查找出当前服务器有多少个活动的SYNC_REC连接,正常的时候这个值应该很小,
  最好是小于5,
当收到Dos攻击或邮件轰炸时这个值相当高.尽管如此这个值和系统有很大的关系,有的服务器值很高也是正常现象
```
  netstat -aulntp | grep SYN_REC | wc -l
```
* 列出所有连接过的IP地址
```
  netstat -n -p | grep SYN_REC | sort -u
```
* 列出所有发送SYN_REC连接节点的IP地址
```
  netstat -n -p | grep SYN_REC | awk '{print $5}' | awk -F: '{print $1}'
```
* 使用netstat命令计算每个主机连接到本机的连接数
```
  netstat -ntu | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -n
```
* 列出所有连接到本机的UDP或者TCP连接的IP数量
```
  netstat -anp |grep 'tcp|udp' | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -n
```
* 检查 ESTABLISHED 连接并且列出每个IP地址的连接数量
```
  netstat -ntu | grep ESTAB | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -nr
```
* 列出所有连接到本机80端口的IP地址和其连接数
```
  netstat -plan|grep :80|awk {'print $5'}|cut -d: -f 1|sort|uniq -c|sort -nk 1
```
