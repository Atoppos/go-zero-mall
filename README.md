# go-zero-mall

---


- 使用微服务开发，api（http）+rpc（grpc），api充当聚合，复杂逻辑以及调用其他服务写在rpc中
- 服务发现与注册中心使用Consul，Consul启用ACL访问限制
- 链路追踪使用Jaeger，服务监控使用Prometheus，均为go-zero原生支持
- 分布式事务使用DTM，具体使用SAGA模式，go-zero无缝接入DTM

#### 技术栈

---

- go-zero
- Consul
- Mysql
- Redis
- Prometheus
- Jaeger
- Grafana
- DTM

#### 目录结构

---

- common [通用组件]
  - cryptx [密码加盐]
  - jwtx [JWT令牌]
- service [所有业务代码，每个服务下分为api、model、rpc]
  - order
  - pay
  - product
  - user
- prometheus.yml [Prometheus配置文件]

## 注意⚠️

---

因为Consul启用了ACL访问限制，但通过DTM客户端向Consul调用其他服务时，目前可能无法携带Token导致调用服务没有权限，已经在DTM交流群交流过这个问题，后续可能会得到解决
