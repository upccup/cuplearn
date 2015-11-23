## 使用pprof 检测go程序运行资源消耗

### 使用 *net/http/pprof* 包监控
 * 在import 中加入 _  "net/http/pprof"
 * 在程序启动之时启动一个独立的http server 
	 ```
	 go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()

	 ```
* 启动程序, 在浏览器中访问 *http://localhost:6060/debug/pprof/* 可以观察到程序的资源消耗情况  

### 生成CPU状态分析图
```
	安装 graphviz (mac 安装 brew install Graphviz)
	运行 go tool pprof  http://localhost:6060/debug/pprof/heap
	运行命令 web 即可生成svg 文件
	在浏览器中打开 svg 文件即可看淡CPU 状态分析图
```
[参考文档](http://blog.golang.org/profiling-go-programs)