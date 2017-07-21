### docker二层网络原理简单说明

#### 知识准备

1. **veth pair**:  Virtual Ethernet Pair简称veth pair, 是一对虚拟网卡，从一张veth网卡发出的数据包可以直接到达它的peer veth,两者之间存在着虚拟链路。使用环境--两个网桥,一个是Linux内核网桥br1,另一个是ovs网桥 br-eth1,现在想把两个网桥连接起来,就可以用veth pair.

2. **Linux bridge**: Linux内核通过一个虚拟的网桥设备来实现桥接的，这个设备可以绑定若干个以太网接口设备，从而将它们桥接起来.网桥设备br0绑定了eth0和eth1。对于网络协议栈的上层来说，只看得到br0，因为桥接是在数据链路层实现的，上层不需要关心桥接的细节。于是协议栈上层需要发送的报文被送到br0，网桥设备的处理代码再来判断报文该被转发到eth0或是eth1，或者两者皆是；反过来，从eth0或从eth1接收到的报文被提交给网桥的处理代码，在这里会判断报文该转发、丢弃、或提交到协议栈上层。
而有时候eth0、eth1也可能会作为报文的源地址或目的地址，直接参与报文的发送与接收（从而绕过网桥）。

 *参考*: [ Linux内核bridge浅析 ](http://blog.csdn.net/h_cszc/article/details/7742955), [Linux下的虚拟Bridge实现](http://www.cnblogs.com/zmkeil/archive/2013/04/21/3034733.html)
 
 3. **docker 网络模式**: 使用docker run创建Docker容器时，可以用--net选项指定容器的网络模式，Docker有以下4种网络模式：
      * host模式，使用--net=host指定。
      * container模式，使用--net=container:NAME_or_ID指定。
      * none模式，使用--net=none指定。
      * bridge模式，使用--net=bridge指定，默认设置。
    详细介绍参考: [Docker网络详解及pipework源码解读与实践](http://www.infoq.com/cn/articles/docker-network-and-pipework-open-source-explanation-practice/)


  docker bridge 网络模式通过 **veth pair**创建两块虚拟网卡, 一块网卡放到容器中一块放在主机上, 以此来解决容器和主机通信. 但是如果要让容器和宿主机以外的机器通信我们可以使用**Linux bridge** 将**veth pair**主机一端的网卡桥接到主机的网卡上, 这样容器就可以通过主机的网卡与外部的主机通信. 操作步骤如下:
  
  1. 创建自定义的docker network:
  
     我们首先需要创建一个br0自定义网桥，这个网桥并不是通过系统命令手动建立的原始Linux网桥，而是通过Docker的create network命令来建立的自定义网桥，这样避免了一个很重要的问题就是我们可以通过设置DefaultGatewayIPv4参数来设置容器的默认路由，这个解决了原始Linux自建网桥不能解决的问题. 用Docker创建网络时我们可以通过设置subnet参数来设置子网IP范围，默认我们可以把整个网段给这个子网，后面可以用ipam driver（地址管理插件）来进行控制。还有一个参数gateway是用来设置br0自定义网桥地址的，其实也就是你这台宿主机的地址啦。

      ```
      docker network create 
        --opt=com.docker.network.bridge.enable_icc=true
        --opt=com.docker.network.bridge.enable_ip_masquerade=false
        --opt=com.docker.network.bridge.host_binding_ipv4=0.0.0.0
        --opt=com.docker.network.bridge.name=br0
        --opt=com.docker.network.driver.mtu=1500
        --ipam-driver=test
        --subnet=容器IP的子网范围，例:192.168.0.0/24
        --gateway=br0网桥使用的IP,也就是宿主机的地址，例:192.168.0.1
        --aux-address=DefaultGatewayIPv4=容器使用的网关地址 例:192.168.223.2
        mynet
      ```
      
      参数说明:
      
                   选项                  |     等同   |        描述       
      :---------------------------------:|:----------:|-----------------------------:
      com.docker.network.bridge.name    |      -     | 创建Linux bridge使用的bridge名称
      com.docker.network.bridge.enable_ip_masquerade | –ip-masq |    启用IP伪装
      com.docker.network.bridge.enable_icc |  –icc   |  	启用或禁用容器间连接
      com.docker.network.bridge.host_binding_ipv4 | –ip  | 绑定容器端口时默认绑定的IP
      com.docker.network.driver.mtu  |    	–mtu  |  	设置容器网络MTU
      
      注: 这些选项有与docker daemon选项等同的选项
       
  
  2. 桥接:
 
     通过Docker命令去创建一个自定义的网络起名为“mynet”，同时会产生一个网桥br0，之后通过更改网络配置文件（在/etc/sysconfig/network-scripts/下ifcfg-br0、ifcfg-默认网络接口名）将默认网络接口桥接到br0上，重启网络后，桥接网络就会生效。Docker默认在每次启动容器时都会将容器内的默认网卡桥接到br0上，而且宿主机的物理网卡也同样桥接到了br0上了。其实桥接的原理就好像是一台交换机，Docker 容器和宿主机物理网络接口都是服务器，通过veth pair这个网络设备像一根网线插到交换机上。至此，所有的容器网络已经在同一个网络上可以通信了，每一个Docker容器就好比是一台独立的虚拟机，拥有和宿主机同一网段的IP，可以实现跨主机访问了。

最后通过 docker run 命令指定 --net=mynet 启动容器
```
docker run -d --name=tomcat --net=mynet tomcat
```

  所有操作执行完成后主机上的网络环境如下:
```
[root@HC-25-60-39]# cat /etc/sysconfig/network-scripts/ifcfg-bond0
    DEVICE=bond0
    BOOTPROTO=none
    BONDING_OPTS="miimon=100 mode=4 xmit_hash_policy=layer2+3"
    TYPE=Ethernet
    IPADDR=172.25.60.39
    NETMASK=255.255.252.0
    ONBOOT=yes
    
[root@HC-25-60-39]# cat /etc/sysconfig/network-scripts/ifcfg-br0
    DEVICE=br0
    TYPE=Bridge
    BOOTPROTO=static
    IPADDR=172.25.60.39
    GATEWAY=172.25.63.254
    NETMASK=255.255.0.0
    ONBOOT=yes
    NOZEROCONF=yes
    IPV6INIT=no
    NM_CONTROLLED=no
```
配置中将 网卡bond0 桥接到 br0
```
[root@HC-25-60-39]# ifconfig
    bond0     Link encap:Ethernet  HWaddr F8:BC:12:34:CF:C4
              UP BROADCAST RUNNING MASTER MULTICAST  MTU:1500  Metric:1
              RX packets:3100686 errors:0 dropped:0 overruns:0 frame:0
              TX packets:1457594 errors:0 dropped:0 overruns:0 carrier:0
              collisions:0 txqueuelen:0
              RX bytes:325907752 (310.8 MiB)  TX bytes:196972999 (187.8 MiB)
    
    br0       Link encap:Ethernet  HWaddr EA:30:7F:E3:81:1F
              inet addr:172.25.60.39  Bcast:172.25.255.255  Mask:255.255.0.0
              UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
              RX packets:2947365 errors:0 dropped:0 overruns:0 frame:0
              TX packets:1353316 errors:0 dropped:0 overruns:0 carrier:0
              collisions:0 txqueuelen:0
              RX bytes:253618014 (241.8 MiB)  TX bytes:182539913 (174.0 MiB)
    veth2698a00 Link encap:Ethernet  HWaddr EA:30:7F:E3:81:1F
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1609 errors:0 dropped:0 overruns:0 frame:0
          TX packets:943787 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:133061 (129.9 KiB)  TX bytes:79482685 (75.8 MiB)
```
配置完成之后的网卡信息, veth2698a00 为 docker 通过 **veth pair** 创建的虚拟网卡在主机上的一端.

通过 brctl 查看桥接情况
```
[root@HC-25-60-39]# brctl show
bridge name     bridge id               STP enabled     interfaces
br0             8000.ea307fe3811f       no              bond0
                                                        veth2698a00
```

可以看到 br0 同事桥接了 bond0 和 veth2698a00

查看容器内的网络情况
```
[root@HC-25-60-39]# docker exec -it tomcat ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:19:3C:BD
          inet addr:172.25.60.189  Bcast:0.0.0.0  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:961500 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1609 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:80890720 (77.1 MiB)  TX bytes:133061 (129.9 KiB)

lo        Link encap:Local Loopback
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:16436  Metric:1
          RX packets:36 errors:0 dropped:0 overruns:0 frame:0
          TX packets:36 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0
          RX bytes:1968 (1.9 KiB)  TX bytes:1968 (1.9 KiB)
```

可以看到 eth0 上绑定的IP为 IPAM 分配的IP, 而 eth0 则是**veth pair** 在容器内的另一端, 通过如下命令可以确定  eth0 和 veth2698a00 互为 peers

在主机上执行:
```
[root@HC-25-60-39]# ip link list
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 16436 qdisc noqueue state UNKNOWN
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: em1: <BROADCAST,MULTICAST,SLAVE,UP,LOWER_UP> mtu 1500 qdisc mq master bond0 state UP qlen 1000
    link/ether f8:bc:12:34:cf:c4 brd ff:ff:ff:ff:ff:ff
.......
8: bond0: <BROADCAST,MULTICAST,MASTER,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP
    link/ether f8:bc:12:34:cf:c4 brd ff:ff:ff:ff:ff:ff
9: br0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN
    link/ether ea:30:7f:e3:81:1f brd ff:ff:ff:ff:ff:ff
10: virbr0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN
    link/ether 52:54:00:7e:c7:5f brd ff:ff:ff:ff:ff:ff
11: virbr0-nic: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN qlen 500
    link/ether 52:54:00:7e:c7:5f brd ff:ff:ff:ff:ff:ff
12: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN
    link/ether 4a:50:20:92:23:48 brd ff:ff:ff:ff:ff:ff
18: veth2698a00: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP
    link/ether ea:30:7f:e3:81:1f brd ff:ff:ff:ff:ff:ff
    
[root@HC-25-60-39]# ethtool -S veth2698a00
NIC statistics:
     peer_ifindex: 17
```

可以看到 和 *veth2698a00* 匹配的设备序号为 17, 接下来查看容器中的网卡设备
```
sh-4.1# ip link list
19: lo: <LOOPBACK,UP,LOWER_UP> mtu 16436 qdisc noqueue state UNKNOWN
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
17: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP
    link/ether 02:42:ac:19:3c:bd brd ff:ff:ff:ff:ff:ff
```

可以看到 eth0 就是 序号为17的设备, 所以可以肯定 容器中的 eht0 和主机上的 veth2698a00 互为 peers


 这样容器访问宿主机之外的主机的方式就清楚了"
docker eth0 --**veth pair**--> host veth2698a00 ---**Linux bridge**----> host br0 ----------> vlan
