### Docker网络插件

### CNM 介绍

围绕着docker的网络目前有两种比较主流的声音, docker 主导的 Container network model(CNM) 和社区主导的 Container network interface(CNI). CNI提供了一种linux的应用容器的插件化网络解决方案.最初由rkt Networking Proposal发展而来.也就是说CNI并不完全针对docker的容器, 而是提供一种普适的容器网络解决方案. 因此它的模型只涉及两个 概念:
* 容器(container): 容器是拥有独立linux网络命名空间的独立单元. 比如rkt/docker创建出来的容器
* 网络(network): 网络指代了可以相互联系的一组实体.这些实体拥有各自独立唯一的ip.这些实体可以使容器,虚拟机,物理机或者其他网络设备.

下面主要介绍一下CNM

相较于CNI,CNM是docker公司力推的网络模型.其主要包括一下概念:
* Sandbox: sandbox包含了一个容器的网络栈.包含了管理容器的网卡.路由表以及DNS设置.一种Sandbox的实现是通过linux的网络命名空间.一个FreeBSD Jail或者其他类似的概念. 一个Sandbox可以包含多个Endpoint
* Endpoint: Endpoint将Sandbox连接到network上. 一个Endpoint的实现可以通过veth pair, Open vSwitch internal port或者其他的方式. 一个Endpoint只能属于一个network, 也只能属于一Sandbox
* Network: network是由一组可以相互通信的Endpoint组成.一个network的实现可以使linux bridge, vlan或者其他方式.一个network钟可以包含很多歌Endpoint.


CNM的接口相较于CNI较为复杂.其提供了remote plugin的方式,进行插件化开发.remote plugin相较于CNI的命令行是通过http请求进行的,更加友好. remote plugin监听一个端口, docker daemon直接通过这个端口与remote plugin进行交互. 下面介绍在进行docker操作时, docker daemon是如何通CNM插件进行交互的.

* Create Network: 这一系列调用发生在使用 docker network create的过程中:
```
/IpamDriver.RequestPool: 创建subnetpool用于分配IP
/IpamDriver.RequestAddress: 为gateway获取IP
/NetworkDriver.CreateNetwork: 创建neutron network和subnet
```
* Create Container: 这一系列调用发生在使用 docker run的过程中, 当然也可以通过 docker network connect 出发
```
/IpamDriver.RequestAddress: 为容器获取IP
/NetworkDriver.CreateEndpoint: 创建neutron port
/NetworkDriver.Join: 为容器和port绑定
/NetworkDriver.ProgramExternalConnectivity:
/NetworkDriver.EndpointOperInfo
```
* Delete Container: 这一系列调用发生在docker delete 删除一个容器的过程中,当然也可以由docker network disconnect 触发
```
/NetworkDriver.RevokeExternalConnectivity
/NetworkDriver.Leave: 容器和port解绑
/NetworkDriver.DeleteEndpoint
/IpamDriver.ReleaseAddress: 删除port并释放IP
```
* Delete Network: 这一系列调用发生在使用 docker network delete的过程中
```
/NetworkDriver.DeleteNetwork: 删除network
/IpamDriver.ReleaseAddress: 释放gateway的IP
/IpamDriver.ReleasePool: 删除subnetpool
```

### IPAM driver

* 介绍: 在docker 网络中, CNM模块通过IPAM(IP address management) driver管理IP地址的分配. Libnetwork内含一个默认的IPAM    driver,同时它也允许动态的增加第三方IPAM driver.用户在创建网络时可以指定IPAM driver.

* remote IPAM driver: 在注册远程network驱动时. libnetwork使用init函数初始化ipams.remote包.libnetwork传递 ipamapi.Callback类型对象给Init()作为参数.该类对象实现了注册IPAM driver RegisterIpamDriver(). remote driver 包使用这个接口去注册带有libnetwork网络控制器的remote driver. remote driver通过Docker 插件包注册与libnetwork通信. ipams.remote提供remote driver 进程的代理

* 数据存储需求: remote driver需要自己管理自己的数据

* Ipam Contract: remote IPAM driver 必须能处理一下请求(实现这些接口)
```
GetDefaultAddressSpaces
RequestPool
ReleasePool
Request address
Release address
```

* IPAM Configuration and flow: 在创建网络的时候，用户可以通过NetworkOptionIpam()函数来提供IPAM相关的配置。
```
func NetworkOptionIpam(ipamDriver string, addrSpace string, ipV4 []*IpamConf, ipV6 []*IpamConf, opts map[string]string) NetworkOption
```
调用者必须提供IPAM驱动的名字，可选提供地址空间、IPv4和IPv6的一组IpamConf结构体。IPAM驱动名称必须提供否则网络创建会失败。
在配置列表中每一个元素都有如下的结构：
```
// IpamConf contains all the ipam related configurations for a network
type IpamConf struct {
	// The master address pool for containers and network interfaces
	PreferredPool string
	// A subset of the master pool. If specified,
	// this becomes the container pool
	SubPool string
	// Input options for IPAM Driver (optional)
	Options map[string]string
	// Preferred Network Gateway address (optional)
	Gateway string
	// Auxiliary addresses for network driver. Must be within the master pool.
	// libnetwork will reserve them if they fall into the container pool
	AuxAddresses map[string]string
}
```

在创建Network时，libnetwork会遍历IpamConf列表并执行下面的请求：
1. 通过RequestPool()请求地址池

2. 通过RequestAddress()请求指定的网关地址，如未指定则选择任意可用地址作为网关。

3. 通过 RequestAddress()请求指定的附加地址。

如果IPv4配置是空的，libnetwork会自动增加一个空的IpamConf结构体。使libnetwork从指定的地址空间中向IPAM驱动请求一个IPv4地址池，如果未指定地址空间则使用IPAM驱动的默认地址空间。如果IPAM驱动无法提供地址池那么将无法创建网络。如果IPv6配置是空的，libnetwork不会做任何操作。在上面点执行过程中，从IPAM驱动传回的数据将会保存在network的IpamInfo结构体中。

创建Endpoint的时候，libnetwork会遍历配置列表并执行下面的操作：
1. 从IPv4池中请求一个IPv4地址，并把它配置成Endpoint接口的地址，成功则停止遍历。

2. 从IPv6池（若存在）中请求一个IPv6地址，并把它配置成Endpoint接口的地址，成功则停止遍历。

以上任何操作执行失败都会导致端点创建失败。

删除Endpoint时，libnetwork会执行以下操作：

1. 释放Endpoint接口的IPv4地址

2. 释放Endpoint接口的IPv6地址（若存在）

删除Network时，libnetwork会遍历IpamData结构列表并向IPAM驱动执行下面的请求：

1. 通过 ReleaseAddress() 释放网关地址

2. 通过 ReleaseAddress() 释放附加地址

3. 通过 ReleasePool() 释放地址池

下面是这些接口的详细介绍:
* #### GetDefaultAddressSpaces:
  为IPAM返回默认的本地和全局地址空间名。地址空间是一组相互不重叠的地址池，并且与其他地址空间的地址池隔离。换句话说，不同的地址空间中可以存在相同的地址池。
地址空间映射到租户名。在libnetwork中，本地和全局地址空间的含义是本地地址空间不需要与集群同步，而全局地址空间则需要与集群同步。除非指定了IPAM配置，否则libnetwork将从默认本地或者默认全局地址空间中请求一个地址池。例如，如果没有指定配置，libnetwork将会从默认本地地址空间中为bridge网络请求地址池，从默认全局地址空间中为overlay网络请求地址池。在注册期间，远程驱动将接收一个URL为 */IpamDriver.GetDefaultAddressSpaces* 没有payload的POST消息。驱动返回的形式：
```
{
	"LocalDefaultAddressSpace": string
	"GlobalDefaultAddressSpace": string
}
```

* #### RequestPool: 
  这个API是用来向IPAM驱动注册地址池的。多次单独（ Multiple identical calls）调用必须返回相同的结果。IPAM驱动负责维护池的引用计数。
```
RequestPool(addressSpace, pool, subPool string, options map[string]string, v6 bool) (string, *net.IPNet, map[string]string, error)
```
远程驱动将会接收一个URL为 /IpamDriver.RequestPool ，payload为下面的POST消息
```
{
	"AddressSpace": string
	"Pool":         string 
	"SubPool":      string 
	"Options":      map[string]string 
	"V6":           bool 
}
```
其中:
* AddressSpace IP地址空间：一组不重叠的地址池
* Pool CIDR格式的IPv4或者IPv6地址池
* SubPool 可选，地址池的子集,CIDR格式的IP地址范围
* Options IPAM驱动选项
* V6 IPAM自选池是否为IPv6的
AddressSpace是唯一必须指定的。如果没有指定Pool，IPAM驱动自动申请一个默认地址池返回。 若想要IPAM自选IPv6池的话必须设置V6标志为true。Pool 为空但SubPool 不为空的情况是非法的，应该拒绝。

成功的响应是这样的:
```
{
	"PoolID": string
	"Pool":   string
	"Data":   map[string]string
}
```
其中:
* PoolID 是地址池的标识符，相同的池要有相同的ID
* Pool 是CIDR格式的地址池
* Data IPAM驱动为地址池提供的元数据



* #### ReleasePool:
  这个API用于释放之前注册的地址池:
```
ReleasePool(poolID string) error
```
这个API,remote driver会接收到一个POST请求*/IpamDriver.ReleasePool* 请求内容如下:
```
{
	"PoolID": string
}
```
其中:
* PoolID 是地址池的标识符，相同的池要有相同的ID

成功返回空值

* #### RequestAddress:
请求一个IP地址, 同样也是一样POST请求, */IpamDriver.ReleaseAddress* 请求内容为:
```
{
	"PoolID":  string
	"Address": string
	"Options": map[string]string
}
```
其中:
* PoolID 地址池标识符
* Address 所请求的IP地址，格式为A.B.C.D。 若地址不满足请求将会失败。若空，IPAM驱动从地址池中选择一个可用地址。
* Options IPAM驱动选项

成功的请求返回如下内容:
```
{
	"Address": string
	"Data":    map[string]string
}
```
其中:
* Address 指定的CIDR格式的地址 (A.B.C.D/MM)
* Data IPAM驱动指定的元数据

* #### ReleaseAddress:
释放一个IP地址,同样是一个POST请求: */IpamDriver.ReleaseAddress*, 请求内容如下:
```
{
	"PoolID": string
	"Address": string
}
```
其中:
* PoolID 地址池标识符
* Address 将要释放的地址

* #### GetCapabilities:
在驱动注册期间，libnetwork会查询该驱动的能力。不要求驱动必须支持。若驱动不支持，libnetwork会自动增加空的能力集到内部驱动处理中。
* #### Capabilities:


参考文档: [IPAM Driver](https://github.com/docker/libnetwork/blob/master/docs/ipam.md), [Docker网络插件](http://www.dockone.io/article/1306), [docker的网络-Containernetworkinterface(CNI)与Containernetworkmodel(CNM)](http://www.lancoude.com/14004666.html)