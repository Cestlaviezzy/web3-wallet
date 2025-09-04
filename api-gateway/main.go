package main

import (
	"log"
	"net/http"

	"web3-wallet/api-gateway/client"
	"web3-wallet/api-gateway/config"
	"web3-wallet/api-gateway/router"
)

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// 初始化gRPC客户端
	clients, err := client.NewGRPCClients(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to initialize gRPC clients: %v", err)
	}
	defer clients.Close()

	// 设置路由
	r := router.Setup(cfg, clients)

	// 启动服务
	log.Printf("🚀 Web3-Wallet API Gateway Starting...")
	log.Printf("📡 Server binding to: %s", cfg.Server.Address())
	log.Printf("🔧 Gateway ready to process requests!")

	if err := http.ListenAndServe(cfg.Server.Address(), r); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
