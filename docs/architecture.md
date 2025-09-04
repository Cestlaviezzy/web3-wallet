# Web3 Wallet 微服务架构设计文档

## 项目概述

Web3-Wallet 是一个基于微服务架构的多链数字钱包系统，采用 Golang + Gin 框架实现。系统支持多条区块链（ETH/BSC/Tron等），提供完整的钱包管理、交易处理、后台运营等功能。

## 技术栈

- **开发语言**: Golang
- **Web框架**: Gin (REST API)
- **服务通信**: gRPC (服务间通信)
- **消息队列**: Kafka/NATS (异步事件)
- **数据库**: MongoDB/PostgreSQL
- **缓存**: Redis (分布式锁、缓存)
- **日志存储**: MongoDB/ElasticSearch
- **容器化**: Docker + Kubernetes + Helm

## 一、服务架构划分

### 1. API Gateway (api-gateway)
**职责**:
- 提供系统统一入口
- 路由分发到各个微服务
- 统一鉴权认证 (JWT/API Key)
- 限流控制
- 请求日志记录
- 跨域处理

**技术特点**:
- 基于 Gin 框架
- 支持多种认证方式
- 集成限流中间件
- 统一错误处理

### 2. Wallet Service (wallet-service) - 核心服务
**职责**:
- 多链钱包地址生成 (ETH/BSC/Tron/...)
- 账户余额查询
- 交易构建与数字签名
- 交易广播到区块链网络
- 对接 chain-service 获取链上数据

**技术特点**:
- 支持 HD 钱包 (BIP32/BIP39/BIP44)
- 多链地址派生算法
- 安全的私钥管理 (HSM/MPC 推荐)
- 交易构建引擎

### 3. Chain Service (chain-service) - 链交互服务
**职责**:
- 管理区块链节点 RPC 连接
- 区块链实时监听 (新区块、交易确认)
- 交易回执查询与状态跟踪
- 链上数据同步
- 节点健康检查与故障转移

**技术特点**:
- 多节点负载均衡
- WebSocket 实时监听
- 交易池监控
- 自动重试机制

### 4. Transaction Service (transaction-service) - 内部账务服务
**职责**:
- 用户充值检测与入账
- 提现申请处理与执行
- 资金归集 (冷热钱包管理)
- 交易幂等性保证
- 并发冲突防护
- 内部账务记录

**核心设计**:
- **幂等性**: 基于 tx_id 唯一性约束
- **防并发**: 分布式锁 + 数据库事务
- **状态机**: 交易状态流转管理
- **异步处理**: 队列化交易处理

### 5. Admin Service (admin-service) - 后台管理服务
**职责**:
- 运营人员后台界面
- 提现审核工作流
- 风控规则配置
- 黑名单管理
- RBAC 权限控制
- 操作审计日志

**权限模型**:
- 基于角色的访问控制 (RBAC)
- 细粒度权限分配
- 操作日志审计
- 多级审批流程

### 6. Notify Service (notify-service) - 异步通知服务
**职责**:
- 交易状态变更通知
- 充值到账提醒
- 系统异常告警
- 邮件/短信/钉钉推送
- Webhook 回调

**通知渠道**:
- 前端实时推送 (WebSocket)
- 邮件通知
- 短信通知
- 钉钉机器人
- 第三方 Webhook

## 二、服务通信架构

### 1. 前端/后台 ↔ 系统
- **协议**: REST API (HTTP/HTTPS)
- **框架**: Gin Router
- **认证**: JWT Token / API Key
- **数据格式**: JSON
- **文档**: Swagger/OpenAPI

### 2. 服务间通信
- **协议**: gRPC (HTTP/2)
- **优势**: 高性能、强类型、代码生成
- **负载均衡**: 客户端负载均衡
- **服务发现**: Consul/Etcd

### 3. 异步事件通信
- **消息队列**: Kafka / NATS
- **事件类型**:
  - 充值到账事件
  - 提现状态变更
  - 交易确认事件
  - 系统告警事件
- **消费模式**: 发布/订阅模式

## 三、关键技术实现重点

### 1. 幂等性保证
```
实现策略:
- 全局唯一 tx_id 标识每笔交易
- 数据库唯一索引约束
- Redis 分布式锁预防
- 状态机防重复提交
```

### 2. 防并发冲突
```
技术方案:
- Redis 分布式锁 (Redlock 算法)
- 数据库悲观锁 (SELECT FOR UPDATE)
- 事务隔离级别控制
- 乐观锁版本号机制
```

### 3. 安全性设计
```
安全措施:
- 私钥管理: HSM 硬件安全模块 / MPC 多方计算
- 避免明文私钥存储
- API 接口加密传输
- 敏感操作多重签名
- 定期安全审计
```

### 4. 审计日志
```
日志策略:
- 所有后台操作记录到 MongoDB/ES
- 操作人员、时间、IP、操作内容
- 交易全生命周期跟踪
- 异常操作实时告警
- 日志不可篡改性保证
```

### 5. 异步解耦
```
事件驱动流程:
充值检测 → Chain Service 发现 → Kafka 事件
       ↓
Transaction Service 消费 → 账务入账 → Kafka 事件
       ↓  
Notify Service 消费 → 推送通知到前端/用户
```

## 四、项目目录结构

```
web3-wallet/
├── api-gateway/                 # 网关服务
│   ├── main.go
│   ├── router/                 # 路由配置
│   ├── middleware/             # 中间件 (认证、限流、日志)
│   └── config/                 # 配置管理
│
├── services/                   # 微服务集合
│   ├── wallet-service/         # 钱包核心服务
│   │   ├── main.go
│   │   ├── proto/             # gRPC 协议定义
│   │   ├── handler/           # gRPC/HTTP 处理器
│   │   ├── service/           # 业务逻辑层
│   │   ├── model/             # 数据模型
│   │   ├── repository/        # 数据访问层
│   │   └── config/
│   │
│   ├── chain-service/          # 区块链交互服务
│   │   ├── main.go
│   │   ├── proto/
│   │   ├── client/            # RPC 客户端
│   │   ├── listener/          # 区块监听器
│   │   └── config/
│   │
│   ├── transaction-service/    # 交易服务
│   │   ├── main.go
│   │   ├── proto/
│   │   ├── service/           # 交易逻辑
│   │   ├── model/
│   │   └── repository/
│   │
│   ├── admin-service/          # 后台管理服务
│   │   ├── main.go
│   │   ├── router/            # HTTP 路由
│   │   ├── handler/           # HTTP 处理器
│   │   ├── service/           # 管理业务逻辑
│   │   ├── model/
│   │   ├── repository/
│   │   └── config/
│   │
│   └── notify-service/         # 通知服务
│       ├── main.go
│       ├── consumer/          # 消息队列消费者
│       ├── service/           # 通知逻辑
│       └── config/
│
├── pkg/                        # 公共组件库
│   ├── logger/                # 统一日志
│   ├── middleware/            # 通用中间件
│   ├── utils/                 # 工具函数 (加密、签名)
│   ├── mq/                    # 消息队列封装
│   ├── db/                    # 数据库连接池
│   ├── config/                # 配置中心
│   └── auth/                  # 认证授权
│
├── proto/                      # 全局 gRPC 协议
│   ├── wallet.proto
│   ├── chain.proto
│   ├── transaction.proto
│   └── admin.proto
│
├── deployments/                # 部署配置
│   ├── docker/                # Docker 构建
│   ├── k8s/                   # Kubernetes YAML
│   └── helm/                  # Helm Charts
│
└── docs/                       # 文档
    ├── architecture.md         # 架构设计 (本文档)
    ├── api.md                  # API 接口文档
    └── deployment.md           # 部署文档
```

## 五、数据流向图

### 充值流程
```
用户转账到地址 → Chain Service 监听 → 交易确认
    ↓
Transaction Service 处理 → 账务入账 → Kafka 事件
    ↓
Notify Service → 通知用户充值成功
```

### 提现流程
```
用户提现申请 → API Gateway → Wallet Service
    ↓
Transaction Service → 风控检查 → Admin Service 审核
    ↓
审核通过 → Wallet Service 构建交易 → Chain Service 广播
    ↓
交易确认 → 状态更新 → Notify Service 通知
```

## 六、部署策略

### 1. 容器化部署
- **Docker**: 每个服务独立镜像
- **Kubernetes**: 容器编排管理
- **Helm**: 一键部署配置

### 2. 监控告警
- **Prometheus**: 指标收集
- **Grafana**: 监控面板
- **AlertManager**: 告警通知

### 3. 日志管理
- **ELK Stack**: 日志收集分析
- **Jaeger**: 分布式链路追踪

## 七、扩展性设计

### 1. 新增区块链支持
- 在 Chain Service 中添加新链适配器
- 更新 Wallet Service 地址生成逻辑
- 配置新链的 RPC 节点

### 2. 水平扩展
- 无状态服务设计，支持多实例部署
- 数据库读写分离
- 缓存集群化部署

### 3. 插件化架构
- 通知渠道插件化
- 风控规则插件化
- 支付通道插件化

---

## 维护说明

本文档需要随着系统演进持续更新，建议每次架构变更后及时同步文档内容。

**文档维护者**: 开发团队
**最后更新**: 2024年
**版本**: v1.0 