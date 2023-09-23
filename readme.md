
## 监控服务器状态，离线报警，收集服务器信息

### 使用

```bash
go build .
```

 | 参数  |  类型  |  默认值  |                  描述                   |
 | :---: | :----: | :------: | :-------------------------------------: |
 |  -s   |        |  flase   |          开启服务端模式server           |
 |  -u   |        |  false   |         查看服务器状态 getdata          |
 |  -l   | string | ":12345" | 监听端口 listen addr (default ":12345") |
 |  -t   |  int   |    60    |           超时报警 Lost Time for alert           |