package handler

import (
	"encoding/json"
	"net/http"

	"web3-wallet/api-gateway/client"
	"web3-wallet/api-gateway/types"
)

// WalletHandler 钱包处理器
type WalletHandler struct {
	clients *client.GRPCClients
}

// NewWalletHandler 创建钱包处理器
func NewWalletHandler(clients *client.GRPCClients) *WalletHandler {
	return &WalletHandler{
		clients: clients,
	}
}

// SendTransaction 发送交易 (核心功能)
func (h *WalletHandler) SendTransaction(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusOK, types.Error(types.StatusBadRequest, "Invalid JSON payload"))
		return
	}

	// TODO: 调用wallet-service的gRPC接口
	// result, err := h.clients.WalletClient.SendTransaction(ctx, &walletpb.SendTransactionRequest{...})

	// 临时响应 (后续替换为实际gRPC调用)
	response := map[string]interface{}{
		"txHash": "0x1234567890abcdef...",
		"status": "pending",
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// Simulate 交易模拟
func (h *WalletHandler) Simulate(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusOK, types.Error(types.StatusBadRequest, "Invalid JSON payload"))
		return
	}

	// TODO: 调用wallet-service的模拟接口
	response := map[string]interface{}{
		"executed":  true,
		"succeeded": true,
		"gasUsed":   "21000",
		"gasLimit":  "21000",
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// GetNonce 获取钱包nonce
func (h *WalletHandler) GetNonce(w http.ResponseWriter, r *http.Request) {
	walletAddress := extractPathParam(r.URL.Path, "/nonce/")
	if walletAddress == "" {
		writeJSONResponse(w, http.StatusOK, types.Error(types.StatusBadRequest, "Missing wallet_address parameter"))
		return
	}

	// TODO: 调用chain-service获取nonce
	response := map[string]interface{}{
		"nonce": "42",
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// GetMetaNonce 获取元交易nonce
func (h *WalletHandler) GetMetaNonce(w http.ResponseWriter, r *http.Request) {
	walletAddress := extractPathParam(r.URL.Path, "/meta_nonce/")
	if walletAddress == "" {
		writeJSONResponse(w, http.StatusOK, types.Error(types.StatusBadRequest, "Missing wallet_address parameter"))
		return
	}

	// TODO: 调用wallet-service获取元交易nonce
	response := map[string]interface{}{
		"metaNonce": "10",
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// GetFeeTokens 获取费用代币信息
func (h *WalletHandler) GetFeeTokens(w http.ResponseWriter, r *http.Request) {
	// TODO: 调用wallet-service获取支持的费用代币
	response := []map[string]interface{}{
		{
			"tokenAddress": "0x0000000000000000000000000000000000000000",
			"symbol":       "BNB",
			"decimals":     18,
			"price":        "300.00",
		},
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// GetActionPointConfig 获取积分配置
func (h *WalletHandler) GetActionPointConfig(w http.ResponseWriter, r *http.Request) {
	// TODO: 调用相关服务获取积分配置
	response := map[string]interface{}{
		"enabled": true,
		"config":  map[string]interface{}{},
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// CustomAuthSimulate 自定义认证交易模拟
func (h *WalletHandler) CustomAuthSimulate(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusOK, types.Error(types.StatusBadRequest, "Invalid JSON payload"))
		return
	}

	// 获取App-Id头部 (如果需要)
	appId := r.Header.Get("App-Id")
	_ = appId

	// TODO: 调用wallet-service的自定义认证模拟接口
	response := map[string]interface{}{
		"executed":  true,
		"succeeded": true,
		"gasUsed":   "25000",
		"gasLimit":  "30000",
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// CustomAuthSend 自定义认证发送交易
func (h *WalletHandler) CustomAuthSend(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONResponse(w, http.StatusOK, types.Error(types.StatusBadRequest, "Invalid JSON payload"))
		return
	}

	// 获取App-Id头部
	appId := r.Header.Get("App-Id")
	_ = appId

	// TODO: 调用wallet-service的自定义认证发送接口
	response := map[string]interface{}{
		"txHash": "0xabcdef1234567890...",
		"status": "pending",
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}
