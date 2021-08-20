# 配置RabbitMQ

在节点启动的config.toml中添加如下配置
```
[Eth.Downstream]
URI= "amqp://guest:guest@localhost:5672"
Exchange = "test-exchange"
RoutingKey = "test-key"
RetryInterval = 500 #ms
```

# 通过RPC导出Block信息到文件
通过http请求
```
curl --data '{"method":"eth_dumpBlock","params":[{"start": "0x1", "end": "0x11", "file": "/home/op/blocks.json"}],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
```
通过js控制台操作
```
geth attach data/eth.ipc

$ eth.dumpBlock(1, 100, "/home/op/blocks.json")
```
