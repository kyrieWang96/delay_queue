## 介绍
delay queue 是基于 Golang 实现的延时队列。基于Redis Zset, 以时间戳作为Score, 主动轮询小于当前时间的元素。新增延迟类型支持：支持延迟多少秒和延迟到具体时间。
## 安装
````
go get -u github.com/yasin-wu/delay-queue
````
推荐使用go.mod
````
require github.com/yasin-wu/delay-queue v1.1.4
````
