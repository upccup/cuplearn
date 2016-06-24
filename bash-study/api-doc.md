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
    
#### 限制哪些字段由API返回
  API的使用者并不总是需要一个资源的完整表示。选择返回字段的功能由来已久，它使得API使用者能够最小化网络阻塞，并加速他们对API的调用。
  使用一个字段查询参数，它包含一个用逗号隔开的字段列表。
  例如，下列请求获得的信息将刚刚足够展示一个在售票的有序列表:
  ```
    GET /tickets?fields=id,subject,customer_name,updated_at&state=open&sort=-updated_at
  ```

#### 更新和创建应该返回一个资源描述
  一个 PUT, POST 或者 PATCH 调用可能会对指定资源的某些字段造成更改，而这些字段本不在提供的参数之列 (例如: created_at 或 updated_at 这两个时间戳)。 为了防止API使用者为了获取更新后的资源而再次调用该API，应当使API把更新(或创建)后的资源作为response的一部分来返回。
  以一个产生创建活动的 POST 操作为例, 使用一个 [HTTP 201 状态代码](https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html#sec9.5) 然后包含一个 [Location header](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.30) 来指向新生资源的URL。

#### 关于是否应该HATEOAS
  **注**: **HATEOAS** (Hypermedia as the Engine of Application State 超媒体作为应用程序状态引擎) </br>
 对于API消费方是否应该创建链接，或者是否应该将链接提供给API，有许多混杂的观点。RESTful的设计原则指定了[HATEOAS](http://apigee.com/about/blog/technology/hateoas-101-introduction-rest-api-style-video-slides) ，
大致说明了与某个端点的交互应该定义在元数据(metadata)之中，这个元数据与输出结果一同到达，并不基于其他地方的信息。
虽然web逐渐依照HATEOAS类型的原则运作（我们打开一个网站首页并随着我们看到的页面中的链接浏览），我不认为我们已经准备好API的HATEOAS了。
 当浏览一个网站的时候，决定点击哪个链接是运行时做出的。然而，对于API，决定哪个请求被发送是在写API集成代码时做出的，并不是运行时。
这个决定可以移交到运行时吗？当然可以，不过顺着这条路没有太多好处，因为代码仍然不能不中断的处理重大的API变化。
也就是说，我认为HATEOAS做出了承诺，但是还没有准备好迎接它的黄金时间。为了完全实现它的潜能，需要付出更多的努力去定义围绕着这些原则的标准和工具。

  目前而言，最好假定用户已经访问过输出结果中的文档&包含资源标识符，而这些API消费方会在制作链接的时候用到。关注标识符有几个优势——网络中的数据流减少了，
API消费方存储的数据也减少了（因为它们存储的是小的标识符而不是包含标识符的URLs）。

  同样的，在URL中提供本文倡导的版本号，对于在一个很长时间内API消费方存储资源标识符（而不是URLs），它更有意义.
总之，标识符相对版本是稳定的，但是表示这一点的URL却不是的！

#### 只返回JSON
  是时候在API中丢弃XML了。XML冗长，难以解析，很难读，他的数据模型和大部分编程语言的数据模型 不兼容,
而他的可扩展性优势在你的主要需求是必须序列化一个内部数据进行输出展示时变得不相干。这里有张google趋势图，[比较XML API 和 JSON API](http://www.google.com/trends/explore?q=xml+api#q=xml%20api%2C%20json%20api&cmpt=q)的热度

#### 字段名称书写格式的 snake_case vs camelCase
    如果你在使用JSON (JavaScript Object Notation) 作为你的主要表示格式，正确的方法就是遵守JavaScript命名约定——对字段名称使用camelCase！如果你要走用各种语言建设客户端库的路线，最好使用它们惯用的命名约定.
  C# & Java 使用[camelCase](https://en.wikipedia.org/wiki/CamelCase), python & ruby 使用[snake_case](https://en.wikipedia.org/wiki/Snake_case)。
    资料：基于从2010年的[camelCase 和 snake_case的眼动追踪研究 (PDF)](http://ieeexplore.ieee.org/xpl/articleDetails.jsp?tp=&arnumber=5521745)，
  snake_case比驼峰更容易阅读20％！这种阅读上的影响会影响API的可勘探性和文档中的示例。

#### 缺省情况下确保漂亮的打印和支持gzip
    一个提供空白符压缩输出的API，从浏览器中查看结果并不美观。虽然一些有序的查询参数（如 ?pretty=true ）可以提供来使漂亮打印生效，一个默认情况下能进行漂亮打印的API更为平易近人。\
  额外数据传输的成本是微不足道的，尤其是当你比较不执行gzip压缩的成本。
    考虑一些用例：假设分析一个API消费者正在调试并且有自己的代码来打印出从API收到的数据——默认情况下这应是可读的。或者，如果消费者抓住他们的代码生成的URL，并直接从浏览器访问它——默认情况下这应是可读的。\
  这些都是小事情。做好小事情会使一个API能被更愉快地使用！

#### 如何处理额外传输的数据呢？
    让我们看一个实际例子。我从GitHub API上拉取了一些数据，默认这些数据使用了漂亮打印（pretty print）。我也将做一些GZIP压缩后的对比。
    ```
        $ curl https://api.github.com/users/veesahni > with-whitespace.txt
        $ ruby -r json -e 'puts JSON JSON.parse(STDIN.read)' < with-whitespace.txt > without-whitespace.txt
        $ gzip -c with-whitespace.txt > with-whitespace.txt.gz
        $ gzip -c without-whitespace.txt ? without-whitespace.txt.gz
    ```
    输出文件的大小如下：
    ```
        without-whitespace.txt - 1252 bytes
        with-whitespace.txt - 1369 bytes
        without-whitespace.txt.gz - 496 bytes
        with-whitespace.txt.gz - 509 bytes
    ```
  在这个例子中，当未启用GZIP压缩时空格增加了8.5%的额外输出大小，而当启用GZIP压缩时这个比例是2.6%。
另一方面，GZIP压缩节省了60%的带宽。由于漂亮打印的代价相对比较小，最好默认使用漂亮打印，并确保GZIP压缩被支持。
关于这点想了解更多的话，Twitter发现当对他们的 [Streaming API](https://dev.twitter.com/streaming/overview) 开启GZIP支持后可以在某些情况获得 [80%的带宽节省](https://dev.twitter.com/blog/announcing-gzip-compression-streaming-apis) 。
  Stack Exchange甚至强制要求必须对API请求结果使用GZIP压缩（[never return a response that's not compressed](https://api.stackexchange.com/docs/compression)）。 

#### 不要默认使用大括号封装，但要在需要的时候支持
    许多API会像下面这样包裹他们的响应信息:
    ```
        {
          "data" : {
            "id" : 123,
            "name" : "John"
          }
        }
    ```
  有不少这样做的理由 - 更容易附加元数据或者分页信息，一些REST客户端不允许轻易的访问HTTP头信息，并且JSONP请求不能访问HTTP头信息.
  无论怎样，随着迅速被采用的标准，比如[CORS](http://www.w3.org/TR/cors/)和[Link header from RFC 5988](http://tools.ietf.org/html/rfc5988#page-6), 大括号封装开始变得不必要。
  我们应当默认不使用大括号封装，而仅在特殊情况下使用它，从而使我们的API面向未来。
