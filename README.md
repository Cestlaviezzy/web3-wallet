# web3-wallet

wallet-system/
├── api-gateway/                  # 网关服务
│   ├── main.go
│   ├── router/
│   ├── middleware/
│   └── config/
│
├── services/
│   ├── wallet-service/           # 钱包服务
│   │   ├── main.go
│   │   ├── proto/                # gRPC proto 定义
│   │   ├── handler/              # gRPC / HTTP handler
│   │   ├── service/              # 业务逻辑 (地址、余额、签名)
│   │   ├── model/                # 数据库实体
│   │   ├── repository/           # DB 操作
│   │   └── config/
│   │
│   ├── chain-service/            # 链交互服务
│   │   ├── main.go
│   │   ├── proto/
│   │   ├── client/               # RPC 调用节点
│   │   ├── listener/             # 区块监听
│   │   └── config/
│   │
│   ├── transaction-service/      # 交易服务
│   │   ├── main.go
│   │   ├── proto/
│   │   ├── service/              # 交易逻辑 (充值/提现/归集)
│   │   ├── model/
│   │   └── repository/
│   │
│   ├── admin-service/            # 后台管理
│   │   ├── main.go
│   │   ├── router/               # gin 路由
│   │   ├── handler/              # HTTP handler
│   │   ├── service/              # 管理逻辑 (风控、审核)
│   │   ├── model/
│   │   ├── repository/
│   │   └── config/
│   │
│   └── notify-service/           # 异步通知
│       ├── main.go
│       ├── consumer/             # Kafka/NATS 消费
│       ├── service/
│       └── config/
│
├── pkg/                          # 公共包（所有服务复用）
│   ├── logger/                   # 日志
│   ├── middleware/               # 中间件
│   ├── utils/                    # 工具类 (签名、加密)
│   ├── mq/                       # Kafka/NATS 封装
│   ├── db/                       # 数据库封装
│   ├── config/                   # 配置中心
│   └── auth/                     # JWT / API Key
│
├── proto/                        # 全局 gRPC proto 定义
│   ├── wallet.proto
│   ├── chain.proto
│   ├── transaction.proto
│   └── admin.proto
│
├── deployments/                  # 部署相关
│   ├── docker/                   # Dockerfile
│   ├── k8s/                      # K8s yaml
│   └── helm/                     # Helm charts
│
└── docs/                         # 架构文档、API 文档
