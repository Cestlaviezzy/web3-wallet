package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"web3-wallet/api-gateway/client"
	"web3-wallet/api-gateway/types"
)

// CommonHandler 通用处理器
type CommonHandler struct {
	clients *client.GRPCClients
}

// NewCommonHandler 创建通用处理器
func NewCommonHandler(clients *client.GRPCClients) *CommonHandler {
	return &CommonHandler{
		clients: clients,
	}
}

// Health 健康检查
func (h *CommonHandler) Health(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	}

	writeJSONResponse(w, http.StatusOK, response)
}

// Status 获取服务状态
func (h *CommonHandler) Status(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "running",
		"version":   "1.0.0",
		"timestamp": time.Now().Unix(),
		"services": map[string]bool{
			"wallet":      true, // TODO: 实际检查服务状态
			"chain":       true,
			"transaction": true,
			"admin":       true,
			"notify":      true,
		},
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// GetChainID 获取链ID
func (h *CommonHandler) GetChainID(w http.ResponseWriter, r *http.Request) {
	// TODO: 调用chain-service获取当前链ID
	response := map[string]interface{}{
		"chainId": "97", // BSC Testnet
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// writeJSONResponse 写入JSON响应
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// extractPathParam 从URL路径中提取参数
func extractPathParam(path, prefix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	return strings.TrimPrefix(path, prefix)
}
