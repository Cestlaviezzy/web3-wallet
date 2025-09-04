package router

import (
	"log"
	"net/http"

	"web3-wallet/api-gateway/client"
	"web3-wallet/api-gateway/config"
	"web3-wallet/api-gateway/handler"
	"web3-wallet/api-gateway/middleware"
)

// Setup 设置路由
func Setup(cfg *config.Config, clients *client.GRPCClients) http.Handler {
	mux := http.NewServeMux()

	// 创建处理器
	commonHandler := handler.NewCommonHandler(clients)
	walletHandler := handler.NewWalletHandler(clients)
	txHandler := handler.NewTransactionHandler(clients)

	// 健康检查
	mux.HandleFunc("/health", commonHandler.Health)

	// 核心钱包API (兼容MetaSignX-Wallet接口)
	mux.HandleFunc("/send_transaction", walletHandler.SendTransaction)
	mux.HandleFunc("/tx_receipt/", txHandler.GetTxReceipt)
	mux.HandleFunc("/simulate", walletHandler.Simulate)
	mux.HandleFunc("/chain_id", commonHandler.GetChainID)
	mux.HandleFunc("/nonce/", walletHandler.GetNonce)
	mux.HandleFunc("/meta_nonce/", walletHandler.GetMetaNonce)
	mux.HandleFunc("/submitters", txHandler.GetSubmitters)
	mux.HandleFunc("/action_point/config", walletHandler.GetActionPointConfig)
	mux.HandleFunc("/fee/tokens", walletHandler.GetFeeTokens)
	mux.HandleFunc("/status", commonHandler.Status)

	// 自定义认证API v1
	mux.HandleFunc("/api/v1/custom_auth/transaction/simulate", walletHandler.CustomAuthSimulate)
	mux.HandleFunc("/api/v1/custom_auth/transaction/send", walletHandler.CustomAuthSend)

	// 应用中间件
	handler := middleware.CORS(mux)
	handler = middleware.Logging(handler)

	log.Printf("🌐 API Endpoints Registered:")
	log.Printf("   • POST /send_transaction - Send wallet transaction")
	log.Printf("   • GET  /tx_receipt/{tx_hash} - Get transaction receipt")
	log.Printf("   • GET  /nonce/{wallet_address} - Get wallet nonce")
	log.Printf("   • GET  /chain_id - Get current chain ID")
	log.Printf("   • GET  /submitters - Get relayer submitter status")
	log.Printf("   • GET  /status - Get service status")
	log.Printf("   • POST /simulate - Simulate transaction")
	log.Printf("   • GET  /fee/tokens - Get supported fee tokens")
	log.Printf("   • POST /api/v1/custom_auth/transaction/* - Custom auth APIs")
	log.Printf("   • GET  /health - Health check")

	return handler
}
