### Redhat 6.5 部署 etcd集群

由于etcd 的部署需要 systemd, 而Redhat并没有systemd 所以etcd 的 rpm 包需要经过特殊处理. [etcd-rpm-el6](https://gitlab.com/pixdrift/etcd-rpm-el6) 给出了解决方案, 所以选择 etcd-2.2 作为安装版本.

1. 环境准备, etcd 的分布式一致性协议采用的是Raft, 所以集群节点最好是单数个, 高可用集群节点最少是三个,所以需要准备3 台机器这里IP 分别是: 172.25.60.35, 172.25.60.51, 172.25.60.39

2. 下载 rpm 包 [下载地址](http://www.pixeldrift.net/rpm/SRPMS/etcd-2.2.0-2.el6.src.rpm), 然后执行 rpm -ivh 安装

3. 配置etcd 集群这一步至关重要, 默认的配置在/etc/sysconfig/etcd 分别在上面的3台机器的配置文件中写入如下内容

* 172.25.60.50
```
ETCD_OPTIONS="--name etcd0 --data-dir /var/lib/etcd/etcd0  --listen-peer-urls http://0.0.0.0:2380 --listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 --initial-advertise-peer-urls http://172.25.60.50:2380 --initial-cluster etcd0=http://172.25.60.50:2380,etcd1=http://172.25.60.35:2380,etcd2=http://172.25.60.39:2380 --initial-cluster-state new --initial-cluster-token mritd-etcd-cluster --advertise-client-urls http://172.25.60.50:2379,http://172.25.60.50:4001"
```

* 172.25.60.35
```
ETCD_OPTIONS="--name etcd1 --data-dir /var/lib/etcd/etcd1  --listen-peer-urls http://0.0.0.0:2380 --listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 --initial-advertise-peer-urls http://172.25.60.35:2380 --initial-cluster etcd0=http://172.25.60.50:2380,etcd1=http://172.25.60.35:2380,etcd2=http://172.25.60.39:2380 --initial-cluster-state new --initial-cluster-token mritd-etcd-cluster --advertise-client-urls http://172.25.60.35:2379,http://172.25.60.35:4001"
```

* 172.25.60.39
```
ETCD_OPTIONS="--name etcd2 --data-dir /var/lib/etcd/etcd2  --listen-peer-urls http://0.0.0.0:2380 --listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 --initial-advertise-peer-urls http://172.25.60.39:2380 --initial-cluster etcd0=http://172.25.60.50:2380,etcd1=http://172.25.60.35:2380,etcd2=http://172.25.60.39:2380 --initial-cluster-state new --initial-cluster-token mritd-etcd-cluster --advertise-client-urls http://172.25.60.39:2379,http://172.25.60.39:4001"
```

参数说明
* --name 节点名称
* --data-dir 数据存放位置
* --listen-peer-urls  Etcd 实例监听的地址
* --listen-client-urls 客户端监听地址
* --initial-advertise-peer-urls 和其他服务通信地址
* --initial-cluster 初始化集群内节点地址
* --initial-cluster-state 初始化集群状态，new 表示新建
* --initial-cluster-token 初始化集群 token
* --advertise-client-urls 客户端访问地址

集群搭建好后，在任意节点执行 etcdctl member list 可列所有集群节点信息，如下所
```
[root@HC-25-60-39 supdev]# etcdctl member list
6cc7b0cef9656834: name=etcd2 peerURLs=http://172.25.60.39:2380 clientURLs=http://172.25.60.39:2379,http://172.25.60.39:4001
7c901eced3ba137d: name=etcd0 peerURLs=http://172.25.60.50:2380 clientURLs=http://172.25.60.50:2379,http://172.25.60.50:4001
fb66b3822e0257b7: name=etcd1 peerURLs=http://172.25.60.35:2380 clientURLs=http://172.25.60.35:2379,http://172.25.60.35:4001
```

存储测试:
  etcd 是一个分布式的key-value 存储, 在任意一台机器上执行如下命令
  ```
  [root@HC-25-60-39 supdev]# etcdctl set aaaa bbbb
bbbb
[root@HC-25-60-39 supdev]# etcdctl get aaaa
bbbb
  ```
  
  可以看到 在本地设置的key值已经成功, 换到另外两台机器中的任意一台,测试是否能获取相同的值
  ```
  [root@HC-25-60-35 supdev]# etcdctl get aaaa
bbbb
  ```
  可以看到存储已经自动同步到别的节点, 致辞etcd 集群搭建成功