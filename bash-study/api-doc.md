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

##### URLs and actions 怎么选择?
  它们应该是有意义于 API 使用者的名词（不是动词). 虽然内部Model可以简单地映射到资源上，但那不一定是个一对一的映射。这里的关键是不要泄漏与API不相关的实现细节。一些相关的名词可以是 票，用户和小组
  一旦定义好了资源, 需要确定什么样的 actions 应用它们，这些 actions 怎么映射到你的 API 上。RESTful 原则提供了 HTTP methods 映射作为策略来处理 CRUD actions(CRUD是指在做计算处理时的增加(Create)、重新取得数据(Retrieve)、更新(Update)和删除(Delete)几个单词的首字母简写。主要被用在描述软件系统中数据库或者持久层的基本操作功能)，如下：
  * GET /tickets - 获取 tickets 列表
  * GET /tickets/12 - 获取一个单独的 ticket
  * POST /tickets - 创建一个新的 ticket
  * PUT /tickets/12 - 更新 ticket #12
  * PATCH /tickets/12 - 部分更新 ticket #12
  * DELETE /tickets/12 - 删除 ticket #12 </br>
  更多信息可以参考 [Roy Felding 的论文 network based software architectures 的第五章](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)

##### 接入点的名称应该选择单数还是复数?
  这里的适用于 keep-it-simple原则
  虽然你内在的语法知识会告诉你用复数形式描述单一资源实例是错误的，但实用主义的答案是保持URL格式一致并且始终使用复数形式。不用处理各种奇形怪状的复数形式（比如person/people，goose/geese）可以让API消费者的生活更加美好，也让API提供者更容易实现API（因为大多数现代框架天然地将/tickets和/tickets/12放在同一个控制器下处理）

##### 如何处理（资源的）关系呢？
如果关系依托于另外一个资源，Restful原则提供了很好的指导原则。让我们来看一个例子:
  * GET /tickets/12/messages - 获取ticket #12下的消息列表
  * GET /tickets/12/messages/5 - 获取ticket #12下的编号为5的消息
  * POST /tickets/12/messages - 为ticket #12创建一个新消息
  * PUT /tickets/12/messages/5 - 更新ticket #12下的编号为5的消息
  * PATCH /tickets/12/messages/5 - 部分更新ticket #12下的编号为5的消息
  * DELETE /tickets/12/messages/5 - 删除ticket #12下的编号为5的消息 </br>

或者如果某种关系不依赖于资源，那么在资源的输出表示中只包含一个标识符是有意义的。API消费者然后除了请求资源所在的接入点外，还得再请求一次关系所在的接入点。但是如果一般情况关系和资源一起被请求，API可以提供自动嵌套关系表示到资源表示中，这样可以防止两次请求API  

##### 如果Action不符合CRUD操作那该怎么办？
  1. 重新构造这个Action，使得它像一个资源的field（我理解为部分域或者部分字段）。这种方法在Action不包含参数的情况下可以奏效。例如一个有效的action可以映射成布尔类型field，并且可以通过PATCH更新资源。
  2. 利用RESTful原则像处理子资源一样处理它。例如，Github的API让你通过PUT /gists/:id/star 来 star a gist ，而通过DELETE /gists/:id/star来进行 unstar 。
  3. 有时候你实在是没有办法将Action映射到任何有意义的RESTful结构。例如，多资源搜索没办法真正地映射到任何一个资源接入点。这种情况，/search 将非常有意义，虽然它不是一个名词。这样做没有问题 - 你只需要从API消费者的角度做正确的事，并确保所做的一切都用文档清晰记录下来了以避免（API消费者的）困惑

#### 关于API 文档
  API的好坏关键看其文档的好坏. 好的API的说明文档应该很容易就被找到，并能公开访问。在尝试任何整合工作前大部分开发者会先查看其文档。当文档被藏于一个PDF之中或要求必须登记信息时，将很难被找到也很难搜索到
  好的文档须提供从请求到响应整个循环的示例。最好的是，请求应该是可粘贴的例子，要么是可以贴到浏览器的链接，要么是可以贴到终端里的curl示例 。 [GitHub](https://developer.github.com/v3/gists/#list-gists) 和 [Stripe](https://stripe.com/docs/api) 在这方面做的非常出色, 可以参考.
  一旦你发布一个公开的API，你必须承诺"在没有通告的前提下，不会更改API的功能" .对于外部可见API的更新，文档必须包含任何将废弃的API的时间表和详情。应该通过博客(更新日志)或者邮件列表送达更新说明(最好两者都通知)。

#### 版本控制
  必须对API进行版本控制。版本控制可以快速迭代并避免无效的请求访问已更新的接入点。它也有助于帮助平滑过渡任何大范围的API版本变迁，这样就可以继续支持旧版本API。
  关于API的版本是否应该包含在URL或者请求头中 莫衷一是。从学术派的角度来讲，它应该出现在请求头中。然而版本信息出现在URL中必须保证不同版本资源的浏览器可浏览性（browser explorability).
  API不可能完全稳定。变更不可避免，重要的是变更是如何被控制的。维护良好的文档、公布未来数月的deprecation计划，这些对于很多API来说都是一些可行的举措。它归根结底是看对于业界和API的潜在消费者是否合理。</br>
  参考文档: [API的版本是否应该包含在URL或者请求头中](http://stackoverflow.com/questions/389169/best-practices-for-api-versioning) </br>
            [approach that Stripe has taken to API versioning](https://stripe.com/docs/api#versioning)

#### 结果过滤，排序和搜索
  最好是尽量保持基本资源URL的简洁性。 复杂结果过滤器、排序需求和高级搜索 (当限定在单一类型的资源时) ，都能够作为在基本URL之上的查询参数来轻松实现。下面让我们更详细的看一下:

* 过滤:
    对每一个字段使用一个唯一查询参数，就可以实现过滤。 例如，当通过“/tickets”终端来请求一个票据列表时，你可能想要限定只要那些在售的票。这可以通过一个像 GET /tickets?state=open 这样的请求来实现。这里“state”是一个实现了过滤功能的查询参数。

* 排序:
    跟过滤类似, 一个泛型参数排序可以被用来描述排序的规则. 为适应复杂排序需求，让排序参数采取逗号分隔的字段列表的形式，每一个字段前都可能有一个负号来表示按降序排序。我们看几个例子:
    ```
        GET /tickets?sort=-priority - 获取票据列表，按优先级字段降序排序
        GET /tickets?sort=-priority,created_at - 获取票据列表，按“priority”字段降序排序。在一个特定的优先级内，较早的票排在前面。
    ```
    
* 搜索:
    有时基本的过滤不能满足需求，这时你就需要全文检索的力量。或许你已经在使用  ElasticSearch 或者其它基于  Lucene 的搜索技术。当全文检索被用作获取某种特定资源的资源实例的机制时， 它可以被暴露在API中，作为资源终端的查询参数，我们叫它“q”。搜索类查询应当被直接交给搜索引擎，并且API的产出物应当具有同样的格式，以一个普通列表作为结果。
    把这些组合在一起，我们可以创建以下一些查询:
    ```
        GET /tickets?sort=-updated_at - 获取最近更新的票
        GET /tickets?state=closed&sort=-updated_at - 获取最近更新并且状态为关闭的票。
        GET /tickets?q=return&state=open&sort=-priority,created_at - 获取优先级最高、最先创建的、状态为开放的票，并且票上有 'return' 字样。
    ```
    
* 一般查询的别名:
    为了使普通用户的API使用体验更加愉快， 考虑把条件集合包装进容易访问的RESTful 路径中。比如上面的，最近关闭的票的查询可以被包装成 GET /tickets/recently_closed
    
