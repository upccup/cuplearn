### 问题: 在测试区已启动的容器中 root 用户不能查看 被监听端口的 PID

```
bash-4.1# netstat -lentp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address               Foreign Address             State       User       Inode      PID/Program name
tcp        0      0 0.0.0.0:51899               0.0.0.0:*                   LISTEN      0          93260417   21/sshd
tcp        0      0 0.0.0.0:22                  0.0.0.0:*                   LISTEN      0          93260421   21/sshd
tcp        0      0 :::51899                    :::*                        LISTEN      0          93260419   21/sshd
tcp        0      0 :::8001                     :::*                        LISTEN      500        100132299  -
tcp        0      0 ::ffff:127.0.0.1:8005       :::*                        LISTEN      501        93284460   -
tcp        0      0 ::ffff:127.0.0.1:9001       :::*                        LISTEN      500        100132309  -
tcp        0      0 :::8080                     :::*                        LISTEN      501        93284449   -
tcp        0      0 :::22                       :::*                        LISTEN      0          93260423   21/sshd
```

如上所示在容器里root用户(UserID:0) 的用户执行netstat 查看端口坚挺情况时, 其他用户(UserID != 0) 启动进程的PID 时不显示的

切换到 tomcat 账户执行同样的明亮结果如下
```
bash-4.1# su tomcat
[tomcat@622c3d9ad50d /]$ netstat -lentp
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address               Foreign Address             State       User       Inode      PID/Program name
tcp        0      0 0.0.0.0:51899               0.0.0.0:*                   LISTEN      0          93260417   -
tcp        0      0 0.0.0.0:22                  0.0.0.0:*                   LISTEN      0          93260421   -
tcp        0      0 :::51899                    :::*                        LISTEN      0          93260419   -
tcp        0      0 :::8001                     :::*                        LISTEN      500        100132299  -
tcp        0      0 ::ffff:127.0.0.1:8005       :::*                        LISTEN      501        93284460   106/java
tcp        0      0 ::ffff:127.0.0.1:9001       :::*                        LISTEN      500        100132309  -
tcp        0      0 :::8080                     :::*                        LISTEN      501        93284449   106/java
tcp        0      0 :::22                       :::*                        LISTEN      0          93260423   -
[tomcat@622c3d9ad50d /]$ id
uid=501(tomcat) gid=501(tomcat) groups=501(tomcat)
```

同样的情况 也只能看到自己起得进程的PID

现在的问题就是: 为什么 root 用户不能查看 tomcat用户所起的进程的PID 呢

参考文章[[docker]privileged参数](http://blog.csdn.net/halcyonbaby/article/details/43499409), 在容器启动时如果没有设置 --privileged container内的root只是外部的一个普通用户权限. 并不是真正的超级管理员. 所以如果需要 root 用户能查看其它用户的资源, 则在启动容器时需要添加 --privileged=true

加上 参数只有的执行结果如下
```
docker run -it -p 6098:22 -p 6099:51899 --privileged=true 172.25.46.9:5001/centos6.5-test-vm-app-biao bash

bash-4.1# netstat -lntpe
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address               Foreign Address             State       User       Inode      PID/Program name
tcp        0      0 ::ffff:127.0.0.1:8005       :::*                        LISTEN      501        100844150  50/java
tcp        0      0 :::8080                     :::*                        LISTEN      501        100844139  50/java
bash-4.1# id
uid=0(root) gid=0(root) 组=0(root)
```

