#### kubernetes是什么
  kubernetes(经常缩写成k8s)是由Google设计的一款开源系统, 用于自动部署,扩展和管理容器(包含但不仅限于docker container). 目前kubernetes主要有一下特性:
  * 容器的自动化部署
  * 自动化扩展活缩容
  * 自动化应用/服务升级
  * 容器成组, 对外提供服务,直接负载均衡
  * 服务的健康检查, 自动重启
 
### kubernetes 中的概念

#### Pods
   kubernetes中最基本的管理单位是pod而不是conatainer.Pod代表部署单位：Kubernetes中单个应用程序的实例，它可能由单个容器或少量紧密耦合并共享资源的容器组成。 一个pod中的所有container都必须运行在同一台机器上.共享网络空间(network namespace) 和存储(volume). Pod的context可以理解为多个Linu namespace的联合
   * PID namespace: 同一Pod中的应用可以看到其他进程
   * network namespace:每一个Pod被分配一个唯一的IP地址. Pod中的所有容器共享 network namespace: 包括IP地址和端口. Pod内的容器可以通过localhost互相访问.当Pod中的容器与Pod外部的实体通信时，他们必须协调他们如何使用共享网络资源（如端口)
   * storage namespace: 一个Pod可以指定一组共享存储的 volumes. Pod中的所有容器都可以访问共享的volumes, 允许这个容器共享数据.Volumes还允许Pod中的数据持久化存储,以防止其中一个容器需要重启.
   * IPC namespace: 同一Pod中的应用可以通过VPC或者POSIX进行通信
   * UTS namespace: 同一Pod中的应用共享一个主机名称

  Pods 在 Kubernetes集群中以多种方式使用,包括:
  * Pods 运行单个容器: “one-container-per-Pod” 是最常见的 kubernetes用例; 在这种情况下可以将Pod视为单个容器的包装,而Kubernetes直接管理Pod而不是容器
  * Pods 运行多个需要在一起工作的容器: Pod可以封装由紧密耦合并需要共享资源的多个共址容器组合成的应用.这些共存容器可能形成一个单一的内部服务单元-- 一个容器通过共享存储提供文件, 而通过另一个"边框"容器来更新这个文件. Pod将这些容器和存储资源作为一个可管理的实体进包装.
  
  每个Pod都是运行给定应用程序的单个实例,如果需要水平扩展应用程序,则应该使用多个Pods,每个实例一个Pod, 在Kubernetes中这种模式称为 replication(复制模式).Replicated Pods通常被当成一个组被抽象想的controller来管理. 一个Controller可以创建和管理多个Pod, 处理复制和扩招，并在集群范围内提供自我修复功能.例如，如果节点出现故障，则Controller可以通过在不同的节点上调度相同的替换来自动替换Pod.

  pod设计的有点
  * 透明，Pod中的容器对基础设施可见使的基础设施可以给容器提供服务，例如线程管理和资源监控，这为用户提供很多便利
  * 解耦软件的依赖关系,独立的容器可以独立的进行重建和重新发布, Kubernetes 甚至会在将来支持独立容器的实时更新
  * 易用: 用户不需要运行自己的线程管理器,也不需关注程序的信号以及异常结束码等
  * 高效: 因为基础设置承载了更多责任,所以容器可以更高效
 
  参考 [Pod Overview](https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/), [名词解释 Pods](https://www.kubernetes.org.cn/kubernetes-pod)


#### services
  在Kubernetes中Pods是并不是一直不变的,Replication Controllers动态的创建和销毁Pods(比如扩容或者缩容和滚动更新--RollingUpdate). 虽然每个Pod有自己的IP,但是这个IP并不能依赖于时间稳定.这样就有一个经常会遇到的问题了:如果一些Pods(称之为前端服务)需要访问另外一下服务(称之为后端服务)的功能,如何让前端服务能持续的追踪到后端服务呢? 这就要用到  Service了
  
  Kubernetes Services 是一组逻辑服务和访问策略的抽象概念.Services中包含Pods的集合通常由Label Selector决定(下文中将看到没有Label Selector的Service).
 
 例如: 现在有一个图像处理的后端服务有3个实例正在运行,这些实例是可替代的,前端服务不关心他们使用的那个后端进程.虽然构成后端集合的实际Pods可能会改变,但是前端客户端并不需要注意或者追踪自己后端的列表. Service 的抽象就可以实现这种解耦.
 
 对于Kubernetes原声应用,Kubernetes提供了一个简单的Endpoint API, 当一个服务中的Pods变化时 API 也随之变化. 对非原生应用,Kubernetes提供了一个基于虚拟IP的网桥服务,这个服务会将请求转发到对应的Pod.
 
 **定义 Service**
 
 和Pod一样Service在Kubernetes中也是REST对象,和所有REST对象一样 Service可以通过POST方法在ApiServer中创建.例如: 假设你有一组服务都开放了9376端口,并且均带有一个标签app=MyApp
```
kind: Service
apiVersion: v1
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```
这段配置会创建一个名叫 *my-service* 的服务,这个服务会定位所有标签中包含 *app=MyApp* 的Pod的TCP端口9376. 新建的服务还会配分配一个IP(有时称作: cluster IP),这个IP由服务代理使用. Service的selector将会持续工作来筛选Pod是,并且将结果POST 到一个Endpoints对象, 这个Endpoints对象的名称也是 **my-service**.

要注意的是 Service 能将任意一个入口端口映射到*targetport*. *targetport*默认值等于属性*port*的值,  *targetport*的值还可以是*string*类型的,指向后端Pods中的*ports*属性.分配给该名称的实际端口号可以在每个后端Pod种不同,这为部署和发展Service提供了很大的灵活性.例如你可以在后端的程序的新版本中更改这个端口而不对客户端造成影响

Kubernetes的 Services 支持TCP和UDP协议, 默认是TCP协议

(未完待续....)
  