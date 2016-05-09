### Linux命令笔记

* 查看所有网卡的网卡名称
```
	ifconfig -a | sed 's/[ \t].*//;/^\(lo\|\)$/d'   (Ubuntu 可用 CentOs会多个冒号)
	ls /sys/class/net/ 
```

* 脚本数值比较操作符

  符号  |  说明
  :---:|:-----:
  -eq  |  等于
  -ne  |  不等于
  -gt  |  大于
  -lt  |  小于
  -ge  |  大于等于
  -le  |  小于等于

* set -e : Exit immediately if a command exits with a non-zero status 如果有非0的结果直接退出脚本

* sudo apt-get install rng-tools && sudo rngd -r /dev/urandom
  生成随机数, 帮助gpg生成key

* cd !$ 把上个命令的参数作为cd参数使用
* cd - 返回进入此目录之前所在的目录


### centos 7 firewall 启动和停止命令
#### 启动
* systemctl enable firewalld
* systemctl start firewalld
* systemctl status firewalld   // running
### 停止
* systemctl disable firewalld
* systemctl stop firewalld
* systemctl status firewalld   // not running

