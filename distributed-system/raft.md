 随着云计算领域的快速发展以及云服务器的普及,分布式系统也因为其自身的可扩展性以及高可用性高度契合了目前的发展潮流而被越来越多的应用到各个应用场景中. 但是在分布式系统中各个服务器之间数据的一致性一直是无法绕过的 难题, 所谓的一致性这个概念, 它是指多个服务器在状态达成一致, 但是在分布式系统中, 因为各种意外可能, 有的服务器可能会崩溃或变得不可靠, 他就不能和其他服务器达成一致性状态. 这样就需要一种Consensus协议, 一致性协议是为了确保容错性, 也就是及时系统中只有一两个服务器宕机, 也不会影响其他的服务器正常提供服务. 在过去Paxos一直是分布式协议的标准,但是Paxos难于理解, 更难以实现.来自Stanford的新的分布式协议研究称为Raft, 他是一个为真实世界应用建立的协议,主要注重协议的落地性和 可理解性.
 
下面简单介绍一下Raft是如何实现在多个服务器之间保证数据的一致性的.为了达成一致性这个目标, 首先Raft需要进行选举,在Raft中，任何时候一个服务器可以扮演下面角色之一：1. Leader: 处理所有客户端交互，日志复制等，一般一次只有一个Leader. 2. Follower: 类似选民，完全被动.3 Candidate候选人: 类似Proposer律师，可以被选为一个新的领导人。参选者需要说服大多数选民(服务器)投票给他. 选举的过程大概分为以下几步:
*  任何一个服务器都可以成为一个候选者Candidate，它向其他服务器Follower发出要求选举自己的请求：
*  其他服务器同意了，发出OK。注意如果在这个过程中，有一个Follower当机，没有收到请求选举的要求，因此候选者可以自己选自己，只要达到N/2 + 1 的大多数票，候选人还是可以成为Leader的。
*  这样这个候选者就成为了Leader领导人，它可以向选民也就是Follower们发出指令，比如进行日志复制。
*  如果一旦这个Leader宕机崩溃了，那么Follower中有一个成为候选者，发出邀票选举。
*  Follower同意后，其成为Leader.

选出leader之后,所有的操作都需要在leader上进行, leader把所有的操作下发到集群中所有的其他服务器(follower). follower收到消息,完成操作commit之后需要向leader汇报状态.如果leader处于不可用的状态,则需要重新进行选举.为了以容错方式达成一致性,我们Raft不要求所有的服务器100%都达成一致, 只要超过半数的大多数服务器达成一致就可以了,假设有N台服务器, N/2+1就超过半数,代表大多数了, 也就是说一个3节点的Raft集群允许一个节点宕机, 一个5节点额Raft节点可以允许2个节点宕机. 所以为了更有效率利用服务器,一般Raft集群里服务器的数量一般都是奇数个,建议配置是运行3或5节点的Raft集群, 这样讲最大限度的提高可用性, 而不会大大的牺牲性能.

进入正题,下面介绍一下如何自己动手写一个内置Raft集群的分布式服务.由于是使用go语言开发,所以选用[etcd/raft](https://github.com/coreos/etcd/tree/master/raft).

在Raft集群中最重要的概念就是一个RaftNode,多个互相连通的RaftNode组成了一个RaftCluster.可以通过以下代码快速启动/重启一个RaftNode.
```
n.Config = &raft.Config{
	ID:              0x01,
	ElectionTick:    3,
	HeartbeatTick:   1,
	Storage:         raft.NewMemoryStorage(),
	MaxSizePerMsg:   1024 * 1024,
	MaxInflightMsgs: 256,
}

if oldwal {
	n.raftNode = raft.RestartNode(n.Config)
} else {
	n.raftNode = raft.StartNode(n.Config, []raft.Peer{{ID: 0x02}, {ID: 0x03}})
}
```
需要注意的一点是这种启动方法是在RaftNode启动时已经知道集群未来的规模,已经将集群的其他节点ID写入到了配置当中,如果要实现节点的Auto Join则需要在Start的Peers参数处传入空值. 另外还有一点需要注意的是如果服务之前已经运行了一段时间,我们在启动服务就需要从WAL中读取上一次服务停止时的状态和数据,然后在这些数据的基础上继续运行.

RaftNode启动成功之后,各个节点之间的通信需要借助一个Transport,这里我们还是使用etcd提供的[httptransport](https://github.com/coreos/etcd/tree/master/rafthttp), 当然了这里也可以基于[grpc](https://github.com/grpc/grpc-go)自己实现所有的通信方法.(可以参考[swarmkit](https://github.com/docker/swarmkit/blob/master/api/raft.proto)的实现). 下面是启动transport的代码实现.
```
	n.transport = &rafthttp.Transport{
		ID:          types.ID(n.id),
		ClusterID:   0x1000,
		Raft:        n,
		ServerStats: ss,
		LeaderStats: stats.NewLeaderStats(strconv.Itoa(n.id)),
		ErrorC:      make(chan error),
	}
	n.transport.Start()

	for i := range n.peers {
		if i+1 != n.id {
			n.transport.AddPeer(types.ID(i+1), []string{n.peers[i]})
		}
	}
	
	ln := newStoppableListener(url.Host, n.httpstopC)

	go func() {
		err := (&http.Server{Handler: n.transport.Handler()}).Serve(ln)
	     <-n.httpstopC
	}()
```
上面的代码中我们先构造出一个transport的实例, 然后将集群中其余节点的访问地址添加到transport中,我们需要通过这点地址来进行节点之间的互相通信.最后我们启动一个TCPListener来接收和Handle接收到的所有消息.

到这里一个简单的RaftNode就算启动起来啦,可以按照同样的方式在启动其他两个节点,这个不是简单的3节点的RaftCluster就算是运行起来了.但是只是服务启动成功还远远不够,RaftCluster最主要的功能是帮助我们保持数据的一致性,那我们应该怎么样做呢?

首先我们将数据提交到RaftNode,使用Propose方法将数据提交到leader节点(注意只能讲数据Propose到leader节点, 因为只有leader才有能力让follower复制自己的操作),leader节点接收到数据首先会根据集群的状态判断是否已经能接受数据的提交,leader确定能接收数据后会负责将数据发送到follower.(注:如果是集群节点的改动需要调用ProposeConfChange方法).需要注意的是对于RaftNode节点数据提交是一个异步的过程,我们通过Propose方法往RaftNode中提交数据,而RaftNode则在经过一系列的状态判断从另一个线程中通过 Ready()方法通知.另外集群状态的变化也会通过这个channel来通知.所以当我们Propose数据之后我们不知道数据是否提交成功,所以如果服务的数据如果有高可用性的要求这里可能需要进行一些额外的处理将异步的提交变成同步的.(可以参考[swarmkit ProposeValue](https://github.com/docker/swarmkit/blob/master/manager/state/raft/raft.go#L1154)).

还有一点需要注意一下的就是因为数据提交只有一个Propose接口,所以如果需要对不同的数据进行不同的操作,我们就需要提前定义好需要对哪些数据(比如app, cluster)进行什么样的操作(比如Add, Update, Delete).这种情况下我们就需要现在Propose的对象里面同时加上数据和操作之后在进行序列化.

下面我们着重看下从 Ready()中收到数据之后的处理,先看代码:
```

	for {
		select {
		case <-n.ticker.C():
			n.raftNode.Tick()
		case rd := <-n.raftNode.Ready():
			n.wal.Save(rd.HardState, rd.Entries)

			n.raftStorage.Append(rd.Entries)
			n.transport.Send(rd.Messages)

			if err := n.publishEntries(n.entriesToApply(rd.CommittedEntries)); err != nil {
				log.L.Errorf("raft: store data failed. Error: %s", err.Error())
				continue
			}

			n.maybeTriggerSnapshot()

			if wasLeader && atomic.LoadUint32(&n.signalledLeadership) != 1 {
				if n.caughtUp() {
					atomic.StoreUint32(&n.signalledLeadership, 1)
					n.leadershipBroadcast.Publish(IsLeader)
				}
			}
			n.raftNode.Advance()
		}
	}
```
可以看到这里我们最主要的两个处理一个是检查RaftCluster的状态是否已经改变,如果该节点已经从follower升级成为leader我们需要通知外部的服务这个改变,以便外部服务做出相应的调整,另一个处理就是publishEntries, 其实就是讲rd.CommittedEntries持久化或者存到相应的地方.我们可以认为这些CommittedEntries就是已经被RaftCluster接收了的可靠消息.

关于节点状态的变化,我们需要在外部服务中监听RaftNode的leadershipChange event,由于RaftNode只有在leader上才能Propose数据(相当于写操作),所以cluster中的所有节点地位并不是对等的,有的需要提交数据的功能可能需要等RaftCluster leader election完成后再leader上启动.至于其他的follower节点如果对外想提供和leader一样的服务能力,则可能需要自己实现一个proxy将请求proxy到leader节点, 或者通过grpc来远程调用leader上相应的接口.相关的代码如下所示:
```
for {
	select {
	case leadershipEvent := <-leadershipCh:
		newState := leadershipEvent.(raft.LeadershipState)

		if newState == raft.IsLeader {
			log.G(ctx).Info("Now i become a leader !!!")
			// TODO do something

		} else if newState == raft.IsFollower {
			log.G(ctx).Info("Now i become a follower !!!")
			//TODO do something
		}
	}
}

```

上面关于接受Ready()的消息的处理代码中除了上面提到的两点外剩下的就是关于WAL(Write Ahead Logging 预写式日志多用于实现数据库原子性事务)和snapshot的相关处理了,通过代码可以看到Ready()里收到的每一条消息都会先调用wal.Save来将相关的信息保存到WAL中,而当操作积累到一定的数量时,则会通过saveSnapshot来将目前的全量数据(包括状态和已经接受到的所有数据)保存到snapshot中, 然后调用wal.ReleaseLockTo释放掉已经存入snapshot中的操作.在这一点上和很多数据库实现的WAL原理( WAL机制的原理是：修改并不直接写入到数据库文件中，而是写入到另外一个称为WAL的文件中；如果事务失败，WAL中的记录会被忽略，撤销修改；如果事务成功，它将在随后的某个时间被写回到数据库文件中，提交修改)是一样的.

最后别忘了调用raftNode.Advance().

到了这里一个简单的RaftCluster就算能正常运行起来啦.当然了这还仅仅只是跑起来一个简答的RaftCluster,关于我们自己的服务怎么和这个嵌入的RaftCluster怎么结合,以及leader切换,节点的增删等等还是有不少的问题等着我们去解决.

有关Raft的更多资料可以参考: [Raft动画演示](http://thesecretlivesofdata.com/raft/), [Raft论文](https://ramcloud.stanford.edu/wiki/download/attachments/11370504/raft.pdf)