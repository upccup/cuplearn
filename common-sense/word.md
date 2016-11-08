## 编程词汇惯例整理

#### get, fetch and query
* .fetch() always returns a list, while .get() returns the first result, or None if there are no results. <br>
* The Database class has two methods for retrieving records Query and Fetch. These are pretty much identical except Fetch 
returns a List<> of POCO's whereas Query uses yield return to iterate over the results without loading the whole set into memory

参考资料 <br>
  [get() method on the Query object](https://cloud.google.com/appengine/docs/python/datastore/queryclass?csw=1#Query_get) <br>
  [petapoco document](http://www.toptensoftware.com/petapoco/)
