## docker 一容器一IP方案研究

#### 目标: 给容器分配一个和宿主机IP在同一网段的IP,让容器充当局域网内的虚拟机, 由于目前的docker版本是1.9.1还不支持制动IP参数所以目前无法实现一容器一指定IP

[Shrike](https://github.com/TalkingData/Shrike)是talkingdata开发的Docker扁平二层网络工具附带IPAM功能.详细启动步骤如下

#### 1. 添加host ip range 防止后面分配的容器IP与主机IP冲突
```
./oam-docker-ipam host-range --ip-start=172.25.60.0/22 --ip-end=172.25.60.182/22 --gateway=172.25.63.254
```

#### 2. 添加ip range 指定容器自动分配IP的范围
```
./oam-docker-ipam ip-range --ip-start=172.25.60.187/24 --ip-end=172.25.60.192/24
```

#### 3. 启动docker-ipam server
```
 ./oam-docker-ipam --debug=true --cluster-store=http://172.25.60.39:2379 server
```

#### 4. 创建主机桥接网络和docker桥接网络
```
./oam-docker-ipam --debug=true --cluster-store=http://172.25.60.39:2379 create-network --ip 172.25.60.39

# 这一步会创建出一个主机桥接网络和docker bridge模式的network

[root@HC-25-60-39 oam-docker-ipam]# ifconfig
bond0     Link encap:Ethernet  HWaddr F8:BC:12:34:CF:C4
          UP BROADCAST RUNNING MASTER MULTICAST  MTU:1500  Metric:1
          RX packets:1948552 errors:0 dropped:0 overruns:0 frame:0
          TX packets:931756 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:196990012 (187.8 MiB)  TX bytes:120408842 (114.8 MiB)

br0       Link encap:Ethernet  HWaddr F8:BC:12:34:CF:C4
          inet addr:172.25.60.39  Bcast:172.25.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1875951 errors:0 dropped:0 overruns:0 frame:0
          TX packets:869624 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:157931908 (150.6 MiB)  TX bytes:111097718 (105.9 MiB)

docker0   Link encap:Ethernet  HWaddr 00:00:00:00:00:00
          inet addr:172.17.0.1  Bcast:0.0.0.0  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:0 (0.0 b)  TX bytes:0 (0.0 b)
          
[root@HC-25-60-39 oam-docker-ipam]# docker network ls
NETWORK ID          NAME                DRIVER
8f8842fa49f7        bridge              bridge
a75a83c4d62f        mynet               bridge
c0a576ef017a        none                null
8a4f0b7501f9        host                host
```

#### 5. 启动容器
```
[root@HC-25-60-39 oam-docker-ipam]# docker run -it --net=mynet busybox ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:19:3C:BC
          inet addr:172.25.60.188  Bcast:0.0.0.0  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:19 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1515 (1.4 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:16436  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)
```

可以看到新起的容器已经分配到了之前指定的IP范围内的IP之一


### 问题

1. 由于talkingdata只提供了rel7的安装包所以**oam-docker-ipam**需要自己编译
2. docker1.9.1 默认的go编译版本是go1.4.3但是目前稳定的go版本已经到了1.8+, 由于go1.6+和之前版本的通信问题(详情参考 https://github.com/moby/moby/issues/20865), 所以需要使用go1.4.3重新编译**oam-docker-ipam** 这里会涉及到一些源码的修改
3. docker1.9.1还不支持 --ip 参数所以不能为容器指定IP
4. 由于docker ipam plugin [Docker IPAM extension API](https://github.com/docker/go-plugins-helpers/tree/master/ipam)在容器重启或停止时或释放掉原来的IP,在重新启动时并不会传原来的IP参数过来所以重启容器会导致IP切换