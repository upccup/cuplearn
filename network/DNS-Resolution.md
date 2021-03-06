## DNS 解析流

DNS(Domain Name System)是"域名系统的缩写",是一种组织成域层次结构的计算机和网络服务命名系统，它用于 TCP/IP 网络，它所提供的服务是用来将主机名和域名转换为 IP 地址的工作。俗话说，DNS 就是将网址转化为对外的 IP 地址。

* 第一步：浏览器会检查缓存中有没有这个域名对应的解析过的IP地址，如果有，该解析过程将会结束。浏览器缓存域名也是有限制的，包括缓存的时间、大小，可以通过 TTL 属性来设置。

* 第二步：如果用户的浏览器缓存中没有，操作系统会先检查自己本地的hosts文件是否有这个网址映射关系，如果有，就先调用这个IP地址映射，完成域名解析。

* 第三步：如果hosts里没有这个域名的映射，则查找本地 DNS 解析器缓存，是否有这个网址映射关系，如果有，直接返回，完成域名解析。

* 第四步：如果hosts与本地 DNS 解析器缓存都没有相应的网址映射关系，首先会找TCP/IP参数中设置的首选 DNS 服务器，在此我们叫它本地 DNS 服务器，此服务器收到查询时，如果要查询的域名，包含在本地配置区域资源中，则返回解析结果给客户机，完成域名解析，此解析具有权威性。

* 第五步：如果要查询的域名，不由本地DNS服务器区域解析，但该服务器已缓存了此网址映射关系，则调用这个IP地址映射，完成域名解析，此解析不具有权威性。

* 第六步：如果本地 DNS 服务器本地区域文件与缓存解析都失效，则根据本地 DNS 服务器的设置（是否设置转发器）进行查询，如果未用转发模式，本地 DNS 就把请求发至13台根 DNS，根 DNS 服务器收到请求后会判断这个域名(.com)是谁来授权管理，并会返回一个负责该顶级域名服务器的一个IP。
  
   本地 DNS 服务器收到IP信息后，将会联系负责 .com 域的这台服务器。这台负责 .com 域的服务器收到请求后，如果自己无法解析，它就会找一个管理 .com 域的下一级 DNS 服务器地址给本地 DNS 服务器。当本地 DNS 服务器收到这个地址后，就会找域名服务器，重复上面的动作，进行查询，直至找到域名对应的主机。
   
* 第七步：如果用的是转发模式，此 DNS 服务器就会把请求转发至上一级 DNS 服务器，由上一级服务器进行解析，上一级服务器如果不能解析，或找根 DNS 或把转请求转至上上级，以此循环。不管是本地 DNS 服务器用的是转发，还是根提示，最后都是把结果返回给本地 DNS 服务器，由此 DNS 服务器再返回给客户机。
