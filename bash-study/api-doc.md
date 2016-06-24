## RESTful API 设计最佳实践

### API设计的一些基本原则
1. API是自述的
```
  一个API有自述性，也就是说看到API的URL，就知道这个API是要干嘛；且这个API的返回值中的字段，又能很好的解释其返回的内容。
  另外它应当对开发者友好并且便于在浏览器地址栏中浏览和探索
```
2. API是完备的
```
  这里主要是指在各种输入参数情况下，API的返回都应该是合理的、完全的. 
```
3. API是抽象的
```
  即通过一层抽象，屏蔽掉具体的数据/实现/细节.
```
4. API是兼容及可扩展的
```
  一个API可能需要同时服务于不同的平台：Web、iOS、Android等，也可能需要服务同一个平台的不同版本.
```
5. 其他
```
  另有很多重要特性，比如安全性、Tracking等
``` 

### 使用 RESTful URLs and actions
  REST的关键原则与将你的 API 分割成逻辑资源紧密相关。使用HTTP请求控制这些资源，其中，这些方法（GET, POST, PUT, PATCH, DELETE）具有特殊含义

##### URLs and actions
  它们应该是有意义于 API 使用者的名词（不是动词). 虽然内部Model可以简单地映射到资源上，但那不一定是个一对一的映射。这里的关键是不要泄漏与API不相关的实现细节。一些相关的名词可以是 票，用户和小组
  一旦定义好了资源, 需要确定什么样的 actions 应用它们，这些 actions 怎么映射到你的 API 上。RESTful 原则提供了 HTTP methods 映射作为策略来处理 CRUD actions，如下：
  * GET /tickets - 获取 tickets 列表
  * GET /tickets/12 - 获取一个单独的 ticket
  * POST /tickets - 创建一个新的 ticket
  * PUT /tickets/12 - 更新 ticket #12
  * PATCH /tickets/12 - 部分更新 ticket #12
  * DELETE /tickets/12 - 删除 ticket #12 </br>
  更多信息可以参考 [Roy Felding 的论文 network based software architectures 的第五章](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)

##### 接入点的名称应该选择单数还是复数
  这里的适用于 keep-it-simple原则
  虽然你内在的语法知识会告诉你用复数形式描述单一资源实例是错误的，但实用主义的答案是保持URL格式一致并且始终使用复数形式。不用处理各种奇形怪状的复数形式（比如person/people，goose/geese）可以让API消费者的生活更加美好，也让API提供者更容易实现API（因为大多数现代框架天然地将/tickets和/tickets/12放在同一个控制器下处理）

##### 处理（资源的）关系
如果关系依托于另外一个资源，Restful原则提供了很好的指导原则。让我们来看一个例子:
  * GET /tickets/12/messages - 获取ticket #12下的消息列表
  * GET /tickets/12/messages/5 - 获取ticket #12下的编号为5的消息
  * POST /tickets/12/messages - 为ticket #12创建一个新消息
  * PUT /tickets/12/messages/5 - 更新ticket #12下的编号为5的消息
  * PATCH /tickets/12/messages/5 - 部分更新ticket #12下的编号为5的消息
  * DELETE /tickets/12/messages/5 - 删除ticket #12下的编号为5的消息 </br>
或者如果某种关系不依赖于资源，那么在资源的输出表示中只包含一个标识符是有意义的。API消费者然后除了请求资源所在的接入点外，还得再请求一次关系所在的接入点。但是如果一般情况关系和资源一起被请求，API可以提供自动嵌套关系表示到资源表示中，这样可以防止两次请求API  
