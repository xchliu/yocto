# YoctoDB

一个玩具级数据库，目的在于通过手撕一个数据库，从而更好地理解数据库系统的实现过程和关键特性；为后续做数据库智能优化，或者是数据库研发做准备。

## 目标

ACID + IDSU


## 技术指标
1. 实现数据的存取即可。当然这也是数据库的本质。    
2. 不打算兼容sql语法，解析器实现时间有点久。最终可以接受只有get和put。
3. 采用go语言。原因是我不会，正好学一下。
4. 采用lsm数据结构，原因我也不会，正好倒腾一下。
5. 采用raft实现多副本，没有原因。
6. 不考虑和数据库周边相关特性。如日志，备份，监控，参数 etc
7. 不考虑性能问题，我的pro已经够慢了。
8. 其他没有想到的，均可随时敲定。
9. 为了方便后续继续玩，尽量模块化，方便重构。



### Milestone
2019-06-18  Startup
2019-07-10  实现初版，进行第一轮insert单线程基准测试。


### Support or Contact

