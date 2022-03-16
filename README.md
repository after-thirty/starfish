# starfish

### feature list
- [X] Memory Session Manager
- [X] DB Session Manager (only support mysql) 
- [ ] RAFT Session Manager  
- [X] Metrics Collector
- [X] TM
- [X] RM TCC
- [X] RM AT
- [X] Client merged request
- [ ] Read config from Config Center
- [ ] Unit Test

### mysql driver

mysql driver 集成 starfish 的工作已经完成，该 driver 基于 https://github.com/go-sql-driver/mysql 开发，开发者可以使用该 driver 对接到各种 orm 中，使用更方便。driver 的项目地址：https://github.com/opentrx/mysql 。 参考 demo：https://github.com/opentrx/starfish-go-samples 。

### 运行 TC

+ 编译
```
cd ${projectpath}/cmd/tc
go build
```

+ 将编译好的程序移动到示例代码目录

```
mv cmd ${targetpath}/
cd ${targetpath}
```

+ 启动 TC

```
./cmd start -config ${projectpath}/cmd/profiles/dev/config.yml
```
