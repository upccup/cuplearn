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
