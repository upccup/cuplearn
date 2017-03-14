## 通过JMX和Permethues/jmx_exproter生成Tomcat监控数据

1. 安装Tomcat: 这里可以直接下载Tomcat的镜像测试, 从dockerhub上下载 tomcat:8.0(随机选了一个镜像做测试).
```
docker pull tomcat:8.0
```

2. 修改Tomcat的启动脚本: 因为下载的是Tomcat的镜像所以启动脚本在镜像里,我们先把容器启动起来,然后把镜像给拷出来.
```
docker run -d --net=host tomcat:8.0
docker cp (containerID):/usr/local/tomcat/bin/catalina.sh ./
```
修改 catalina.sh: 在文件头部添加如下内容:
```
CATALINA_OPTS="-Dcom.sun.management.jmxremote
-Dcom.sun.management.jmxremote.authenticate=false
-Dcom.sun.management.jmxremote.ssl=false
-Dcom.sun.management.jmxremote.port=799 #定义jmx监听端口
-Djava.rmi.server.hostname=192.168.88.206（客户端IP）"
```
以上内容用来开启JMX来收集Tomcat的运行时数据
然后将修改后的文件拷回到容器中去
```
docker cp ./catalina.sh ($containerID):/usr/local/tomcat/bin/catalina.sh
```
提交保存修改
```
docker commit ($containerID)
```
重新给镜像打个Tag
```
docker tag ($imagesID)) tomcat-custom:8.0
```
3: 测试JMX 是否启动成功: 可以使用cmdline-jmxclient-0.10.3.jar来测试
```
java -jar /tmp/cmdline-jmxclient-0.10.3.jar - 192.168.88.206:7199 java.lang:type=Memory NonHeapMemoryUsage
```
出现如下内容，则是正常的
```
01/13/2017 06:30:14 +0800 org.archive.jmx.Client NonHeapMemoryUsage:
committed: 24313856
init: 24313856
max: 136314880
used: 17670632
```
4:下载jmx_exporter镜像(也可以自己使用mvn编译)
```
docker pull frontporch/jmx_exporter
```
5:准备jmx_exporter的Tomcat配置文件
```
---
hostPort: 192.168.59.104:7199
lowercaseOutputLabelNames: true
lowercaseOutputName: true
rules:
- pattern: 'Catalina<type=GlobalRequestProcessor, name=\"(\w+-\w+)-(\d+)\"><>(\w+):'
  name: tomcat_$3_total
  labels:
    port: "$2"
    protocol: "$1"
  help: Tomcat global $3
  type: COUNTER
- pattern: 'Catalina<j2eeType=Servlet, WebModule=//([-a-zA-Z0-9+&@#/%?=~_|!:.,;]*[-a-zA-Z0-9+&@#/%=~_|]), name=([-a-zA-Z0-9+/$%~_-|!.]*), J2EEApplication=none, J2EEServer=none><>(requestCount|maxTime|processingTime|errorCount):'
  name: tomcat_servlet_$3_total
  labels:
    module: "$1"
    servlet: "$2"
  help: Tomcat servlet $3 total
  type: COUNTER
- pattern: 'Catalina<type=ThreadPool, name="(\w+-\w+)-(\d+)"><>(currentThreadCount|currentThreadsBusy|keepAliveCount|pollerThreadCount|connectionCount):'
  name: tomcat_threadpool_$3
  labels:
    port: "$2"
    protocol: "$1"
  help: Tomcat threadpool $3
  type: GAUGE
- pattern: 'Catalina<type=Manager, host=([-a-zA-Z0-9+&@#/%?=~_|!:.,;]*[-a-zA-Z0-9+&@#/%=~_|]), context=([-a-zA-Z0-9+/$%~_-|!.]*)><>(processingTime|sessionCounter|rejectedSessions|expiredSessions):'
  name: tomcat_session_$3_total
  labels:
    context: "$2"
    host: "$1"
  help: Tomcat session $3 total
  type: COUNTER
```
在[官方文档](https://github.com/prometheus/jmx_exporter/blob/master/example_configs/tomcat.yml)的基础上加了最上面一行用来配置JMX的地址

6:启动jmx_exporter
```
docker run -it --net=host -v $(pwd)/jmx_exporter.yaml:/etc/jmx_exporter/jmx_exporter.yaml frontporch/jmx_exporter
```
**jmx_exporter.yaml**就是上面编辑的文件
7:访问**http://localhost/metrics**就可以查看相应的数据

参考文档

[Zabbix通过jmx监控Tomcat](http://www.rfyy.net/archives/1961.html)

[JMX Exporter](https://github.com/prometheus/jmx_exporter)

[Prometheus jmx_exporter in a container](https://github.com/frontporch/jmx_exporter)
