# Web3-Wallet API Gateway

API Gateway是Web3钱包系统的统一入口，负责路由转发、认证授权、限流控制等功能。

## 项目结构

```
api-gateway/
├── main.go              # 启动入口
├── config/
│   └── config.go        # 配置管理
├── client/
│   └── grpc.go         # gRPC客户端管理
├── handler/
│   ├── common.go       # 通用处理器
│   ├── wallet.go       # 钱包相关处理器
│   └── transaction.go  # 交易相关处理器
├── middleware/
│   └── middleware.go   # CORS、日志中间件
├── router/
│   └── router.go       # 路由设置
├── types/
│   └── response.go     # 响应类型定义
└── README.md           # 本文档
```

## 功能特点

### ✅ **API兼容性**
- 完全兼容MetaSignX-Wallet的API接口
- 保持相同的请求/响应格式
- 支持自定义认证API

### 🔗 **微服务集成**
- gRPC客户端管理
- 服务发现和负载均衡
- 故障隔离和重试机制

### 🛡️ **中间件支持**
- CORS跨域处理
- 请求日志记录
- 预留认证、限流扩展接口

## 支持的API接口

### 核心钱包API
```
POST /send_transaction              - 发送交易 (核心功能)
GET  /tx_receipt/{tx_hash}         - 查询交易回执
POST /simulate                      - 交易模拟执行
GET  /chain_id                      - 获取链ID
GET  /nonce/{wallet_address}       - 获取钱包nonce
GET  /meta_nonce/{wallet_address}  - 获取元交易nonce
GET  /submitters                    - 查询提交者状态
GET  /action_point/config           - 获取积分配置
GET  /fee/tokens                    - 获取费用代币信息
GET  /status                        - 获取服务状态
```

### 自定义认证API
```
POST /api/v1/custom_auth/transaction/simulate  - 自定义认证交易模拟
POST /api/v1/custom_auth/transaction/send      - 自定义认证发送交易
```

### 系统API
```
GET  /health                        - 健康检查
```

## 环境变量配置

```bash
# 服务器配置
PORT=3050                                    # 监听端口
HOST=0.0.0.0                                # 监听地址

# 后端微服务地址
WALLET_SERVICE_ADDR=localhost:50001          # 钱包服务
CHAIN_SERVICE_ADDR=localhost:50002           # 链交互服务
TRANSACTION_SERVICE_ADDR=localhost:50003     # 交易服务
ADMIN_SERVICE_ADDR=localhost:50004           # 管理服务
NOTIFY_SERVICE_ADDR=localhost:50005          # 通知服务
```

## 快速启动

```bash
# 1. 进入api-gateway目录
cd web3-wallet/api-gateway

# 2. 启动服务
go run main.go
```

启动成功后，你会看到类似输出：
```
🚀 Web3-Wallet API Gateway Starting...
📡 Initializing gRPC clients...
   • Wallet Service: localhost:50001
   • Chain Service: localhost:50002
   • Transaction Service: localhost:50003
   • Admin Service: localhost:50004
   • Notify Service: localhost:50005
✅ gRPC clients initialized successfully
🌐 API Endpoints Registered:
   • POST /send_transaction - Send wallet transaction
   • GET  /tx_receipt/{tx_hash} - Get transaction receipt
   ...
📡 Server binding to: 0.0.0.0:3050
🔧 Gateway ready to process requests!
```

## API测试

### 1. 健康检查
```bash
curl http://localhost:3050/health
```

### 2. 获取服务状态
```bash
curl http://localhost:3050/status
```

### 3. 发送交易 (示例)
```bash
curl -X POST http://localhost:3050/send_transaction \
  -H "Content-Type: application/json" \
  -d '{
    "from": "0x...",
    "to": "0x...",
    "value": "1000000000000000000"
  }'
```

## 开发说明

### 添加新的API接口
1. 在对应的handler文件中添加处理方法
2. 在router/router.go中注册路由
3. 如需调用后端服务，通过gRPC客户端调用

### 添加新的中间件
1. 在middleware/middleware.go中实现中间件函数
2. 在router/router.go中应用中间件

### 响应格式
所有API响应都遵循MetaSignX-Wallet的格式：
```json
{
  "statusCode": 200,
  "message": "200 OK", 
  "data": {...}
}
```

## 后续开发计划

- [ ] 实际的gRPC客户端集成
- [ ] JWT认证中间件
- [ ] 限流中间件
- [ ] 监控指标收集
- [ ] 配置文件支持
- [ ] Docker容器化

## 注意事项

⚠️ **当前版本为开发版本**
- gRPC客户端暂时未实现，返回模拟数据
- 需要等待后端微服务完成后进行实际集成
- 生产环境需要添加认证、限流等安全措施 