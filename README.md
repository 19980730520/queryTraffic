# queryTraffic

通过xray内置的api查询流量统计

半成品, 只能查询当前连接的流量统计

当连接断开后会重置 ( 切换节点, 客户端关闭时 )

x-ui面板每次断开都会将本村连接写入数据库, 实现历史流量统计

使用方法
```sh
./xqt -s 127.0.0.1:62789 -ptn "inbound>>>inbound-48426>>>traffic>>>downlink"
```

-s 指定api入口的地址及端口, 需要在配置文件中开启api, 详见[xray官方文档](https://xtls.github.io/config/api.html#apiobject)


-ptn 指定查询语句, 详见 [xray官方文档](https://xtls.github.io/config/stats.html)


本项目的实际意义是记录一下我与go的第一次接触
