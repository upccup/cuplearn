### Linux命令笔记

* 查看所有网卡的网卡名称
```
	ifconfig -a | sed 's/[ \t].*//;/^\(lo\|\)$/d'   (Ubuntu 可用 CentOs会多个冒号)
	ls /sys/class/net/ 
```