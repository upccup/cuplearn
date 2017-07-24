## ip 命令详解

linux 的 ip 命令和ifconfig类似,但前者功能更强大,并旨在取代后者. ifconfig是net-tools中已经被废弃使用的一个命令,已经很久没有维护了. iproute2套件里提供了许多增强功能的命令, ip 命令就是其中之一.


        net-tools          |      iproute2       
:-------------------------:|:--------------------------------:
       arp-na              |         ip neigh
       ifconfig            |         ip link
      ifconfig -a          |         ip addr show
      ifconfig --help      |         ip help
      ifconfig -s          |         ip -s link        
      ifconfig eth0 up     |    ip link set eth0 up 
      ipmaddr              |         ip maddr
      iptunnel             |         ip tunnel
      netstat              |            ss
      netstat -i           |         ip -s link
      netstat -g           |         ip maddr
      netstat -l           |            ss -l
      netstat -r           |         ip route
      route add            |         ip route add
      route del            |         ip route del
      route -n             |         ip route show
      vconfig              |         ip link
      

* #### 设置和删除IP
给机器设置一个IP地址可以使用如下命令
```
sudo ip addr add 192.168.0.193/24 dev wlan0
```
查看IP是否生效
```
ip addr show wlan0
```
删除IP地址
```
sudo ip addr del 192.168.0.193/24 dev wlan0
```

* #### 列出路由条目
ip命令的路由对象的参数还可以查看网络中的路由参数
```
ip route show
```
如果已经有一个IP地址,需要知道路由包从哪里来,可以使用如下命令 
```
ip route get 10.42.0.47
```

* #### 更改默认路由
如果要更改默认路由,可以使用一下命令
```
sudo ip route add default via 192.168.0.196
```

* #### 显示网络统计数据
使用ip命令可以显示不同网络接口的统计数据
```
[root@HC-25-60-51]# ip  -s link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 16436 qdisc noqueue state UNKNOWN
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    RX: bytes  packets  errors  dropped overrun mcast
    85040489   637680   0       0       0       0
    TX: bytes  packets  errors  dropped carrier collsns
    85040489   637680   0       0       0       0
2: eth2: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN qlen 1000
    link/ether ec:f4:bb:c1:9c:46 brd ff:ff:ff:ff:ff:ff
    RX: bytes  packets  errors  dropped overrun mcast
    0          0        0       0       0       0
    TX: bytes  packets  errors  dropped carrier collsns
    0          0        0       0       0       0
3: bond0: <BROADCAST,MULTICAST,MASTER,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP
    link/ether ec:f4:bb:c1:9c:44 brd ff:ff:ff:ff:ff:ff
    RX: bytes  packets  errors  dropped overrun mcast
    1944556979 59927311 0       0       0       3079577
    TX: bytes  packets  errors  dropped carrier collsns
    2071635785 55135995 0       0       0       0
4: br0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN
    link/ether da:57:76:bb:17:fb brd ff:ff:ff:ff:ff:ff
    RX: bytes  packets  errors  dropped overrun mcast
    1039973681 59184269 0       0       0       2932527
    TX: bytes  packets  errors  dropped carrier collsns
    1904325073 53163962 0       0       0       0
```

当需要获取一个特定网路接口的信息时,在网络接口名字后面添加选项ls即可. 使用多个选项 -s 会给出更详细的信息
```
[root@HC-25-60-51]# ip  -s -s link ls br0
7: br0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN
    link/ether da:57:76:bb:17:fb brd ff:ff:ff:ff:ff:ff
    RX: bytes  packets  errors  dropped overrun mcast
    1042326454 59210911 0       0       0       2933761
    RX errors: length  crc     frame   fifo    missed
               0        0       0       0       0
    TX: bytes  packets  errors  dropped carrier collsns
    1907027596 53186922 0       0       0       0
    TX errors: aborted fifo    window  heartbeat
               0        0       0       0
```

* #### ARP条目
地址解析协议(ARP)用于将一个IP地址转换成它对应的物理地址,也就是通过所说的MAC地址. 使用ip命令的neigh或者neighbour选项,可以查看介入局域网内的设备的MAC地址
```
ip neighbour
```

* #### 监控netlink消息
可以使用ip命令查看netlink消息.monitor选项可以查看网络设备的状态,比如所在局域网的一台电脑根据它的状态可以被分类成REACHABLE或者STALE.
```
[root@HC-25-60-51]# ip monitor all
[NEIGH]172.25.8.15 dev br0 lladdr 70:ba:ef:6d:cf:14 STALE
[NEIGH]172.25.60.35 dev br0  FAILED
[NEIGH]172.25.63.254 dev br0 lladdr 70:ba:ef:6d:cf:14 REACHABLE
[NEIGH]172.25.60.35 dev br0  FAILED
[NEIGH]172.25.60.35 dev br0  FAILED
[NEIGH]172.25.62.188 dev br0 lladdr a0:36:9f:db:fd:80 REACHABLE
```

* #### 激活和停止网络接口
```
sudo ip link set ppp0 down
sudo ip link set ppp0 up
```

* #### 激活和停止网卡设备
```
sudo ip link set dev br0 up
sudo ip link set dev br0 down
```