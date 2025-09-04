package client

import (
	"log"

	"web3-wallet/api-gateway/config"
)

// GRPCClients gRPC服务客户端集合
type GRPCClients struct {
	// TODO: 这里后续会添加实际的gRPC客户端
	// WalletClient  walletpb.WalletServiceClient
	// ChainClient   chainpb.ChainServiceClient
	// TxClient      txpb.TransactionServiceClient
	// AdminClient   adminpb.AdminServiceClient
	// NotifyClient  notifypb.NotifyServiceClient

	config *config.Config
}

// NewGRPCClients 创建gRPC客户端集合
func NewGRPCClients(cfg *config.Config) (*GRPCClients, error) {
	log.Printf("📡 Initializing gRPC clients...")
	log.Printf("   • Wallet Service: %s", cfg.Services.WalletAddr)
	log.Printf("   • Chain Service: %s", cfg.Services.ChainAddr)
	log.Printf("   • Transaction Service: %s", cfg.Services.TxAddr)
	log.Printf("   • Admin Service: %s", cfg.Services.AdminAddr)
	log.Printf("   • Notify Service: %s", cfg.Services.NotifyAddr)

	clients := &GRPCClients{
		config: cfg,
	}

	// TODO: 初始化实际的gRPC连接
	// conn, err := grpc.Dial(cfg.Services.WalletAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	//     return nil, fmt.Errorf("failed to connect wallet service: %w", err)
	// }
	// clients.WalletClient = walletpb.NewWalletServiceClient(conn)

	log.Printf("✅ gRPC clients initialized successfully")
	return clients, nil
}

// Close 关闭所有gRPC连接
func (c *GRPCClients) Close() {
	log.Printf("🔌 Closing gRPC connections...")
	// TODO: 关闭实际的gRPC连接
}

// GetConfig 获取配置
func (c *GRPCClients) GetConfig() *config.Config {
	return c.config
}
